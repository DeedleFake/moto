package moto

import (
	"context"
	"fmt"

	wl "deedles.dev/wl/server"
	"deedles.dev/xsync"
	"golang.org/x/exp/slog"
)

// Display is a Wayland display for clients to connect to. It handles
// running a socket, listening for connections, and dispatching events
// to the right places.
type Display struct {
	server *wl.Server
	queue  xsync.Queue[func()]
	serial uint32
}

func NewDisplay() (*Display, error) {
	var d Display

	server, err := wl.CreateServer()
	if err != nil {
		return nil, fmt.Errorf("create server: %w", err)
	}
	server.Handler = d.client
	d.server = server

	return &d, nil
}

func (d *Display) Close() error {
	d.queue.Stop()
	return nil
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
					slog.Error("client %p event", err, client)
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

type displayListener struct {
	display *Display
	client  *wl.Client
}

func (lis *displayListener) Sync(cb *wl.Callback) {
	cb.Done(lis.display.NextSerial())
}

func (lis *displayListener) GetRegistry(r *wl.Registry) {
	// TODO
}
