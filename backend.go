package moto

import (
	"os"

	wl "deedles.dev/wl/server"
	"deedles.dev/wl/wire"
)

type Backend interface {
}

func AutocreateBackend(display *Display) (Backend, error) {
	display.AddGlobal(wl.ShmInterface, wl.ShmVersion, func(state wire.State, id wire.NewID) {
		shm := wl.BindShm(state, id)
		shm.Listener = &shmListener{
			display: display,
		}
	})

	display.AddGlobal(wl.CompositorInterface, wl.CompositorVersion, func(state wire.State, id wire.NewID) {
		c := wl.BindCompositor(state, id)
		c.Listener = &compositorListener{
			display: display,
		}
	})

	return nil, nil
}

type shmListener struct {
	display *Display
}

func (lis *shmListener) CreatePool(pool *wl.ShmPool, file *os.File, size int32) {
	// TODO
}

type compositorListener struct {
	display *Display
}

func (lis *compositorListener) CreateSurface(s *wl.Surface) {
	// TODO
}

func (lis *compositorListener) CreateRegion(r *wl.Region) {
	// TODO
}
