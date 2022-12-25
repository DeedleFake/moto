package moto

import (
	"context"
	"fmt"
	"log"

	"deedles.dev/moto/internal/set"
	wl "deedles.dev/wl/server"
	"deedles.dev/wl/wire"
	"deedles.dev/xsync"
)

// Display is a Wayland display for clients to connect to. It handles
// running a socket, listening for connections, and dispatching events
// to the right places.
type Display struct {
	server *wl.Server
	queue  xsync.Queue[func()]
	serial uint32

	reg     set.Set[*wl.Registry]
	name    uint32
	globals map[uint32]global
}

func NewDisplay() (*Display, error) {
	d := Display{
		reg:     make(set.Set[*wl.Registry]),
		globals: make(map[uint32]global),
	}

	server, err := wl.CreateServer()
	if err != nil {
		return nil, fmt.Errorf("create server: %w", err)
	}
	server.Handler = d.client
	d.server = server

	return &d, nil
}

// Serial returns the current event serial.
func (d *Display) Serial() uint32 {
	return d.serial
}

// NextSerial increments the internal event serial count and returns
// the old value.
func (d *Display) NextSerial() uint32 {
	s := d.serial
	d.serial++
	return s
}

func (d *Display) client(ctx context.Context, client *wl.Client) {
	log.Printf("client connected: %p", client)
	defer log.Printf("client disconnected: %p", client)

	// TODO: Can this deadlock?
	defer func() { d.queue.Add() <- func() { client.DeleteAll() } }()

	client.Display().Listener = &displayListener{
		display: d,
		client:  client,
	}

	for {
		select {
		case <-ctx.Done():
			return

		case ev, ok := <-client.Events():
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case d.queue.Add() <- func() {
				err := ev()
				if err != nil {
					log.Printf("error in client %p event: %v", client, err)
				}
			}:
			}
		}
	}
}

// Run runs the display's event loop. It blocks until the display
// exits either as a result of ctx being cancelled or the Display
// being closed. The Display should not be reused after this method
// has returned.
func (d *Display) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go d.server.Run(ctx)

	for {
		select {
		case <-ctx.Done():
			return

		case ev, ok := <-d.queue.Get():
			if !ok {
				return
			}

			ev()
		}
	}
}

func (d *Display) addRegistry(reg *wl.Registry) {
	d.reg.Add(reg)
	for name, g := range d.globals {
		reg.Global(name, g.Interface, g.Version)
	}
}

func (d *Display) removeRegistry(reg *wl.Registry) {
	delete(d.reg, reg)
}

func (d *Display) AddGlobal(inter string, version uint32, create func(state wire.State, id wire.NewID)) (name uint32) {
	name = d.name
	d.name++

	d.globals[name] = global{
		Interface: inter,
		Version:   version,
		Create:    create,
	}

	for reg := range d.reg {
		reg.Global(name, inter, version)
	}

	return name
}

func (d *Display) RemoveGlobal(name uint32) {
	delete(d.globals, name)
	for reg := range d.reg {
		reg.GlobalRemove(name)
	}
}

type global struct {
	Interface string
	Version   uint32
	Create    func(state wire.State, id wire.NewID)
}

type displayListener struct {
	display *Display
	client  *wl.Client
}

func (lis *displayListener) Sync(cb *wl.Callback) {
	cb.Done(lis.display.NextSerial())
}

func (lis *displayListener) GetRegistry(reg *wl.Registry) {
	lis.display.addRegistry(reg)

	reg.OnDelete = func() { lis.display.removeRegistry(reg) }
	reg.Listener = &registryListener{
		display: lis.display,
		client:  lis.client,
	}
}

type registryListener struct {
	display *Display
	client  *wl.Client
}

func (lis *registryListener) Bind(name uint32, id wire.NewID) {
	g, ok := lis.display.globals[name]
	if !ok {
		panic("Not implemented.")
	}

	g.Create(lis.client, id)
}
