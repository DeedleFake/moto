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
	return nil, nil
}

type shmListener struct {
	display *Display
}

func (lis *shmListener) CreatePool(pool *wl.ShmPool, file *os.File, size int32) {
	// TODO
}
