package moto

import (
	wl "deedles.dev/wl/server"
	"deedles.dev/wl/wire"
)

type Compositor struct {
	display  *Display
	renderer *Renderer
}

func CreateCompositor(display *Display, renderer *Renderer) *Compositor {
	compositor := Compositor{
		display:  display,
		renderer: renderer,
	}

	display.AddGlobal(wl.CompositorInterface, wl.CompositorVersion, func(state wire.State, id wire.NewID) {
		c := wl.BindCompositor(state, id)
		c.Listener = &compositorListener{
			compositor: &compositor,
		}
	})

	return &compositor
}

type compositorListener struct {
	compositor *Compositor
}

func (lis *compositorListener) CreateSurface(s *wl.Surface) {
	// TODO
}

func (lis *compositorListener) CreateRegion(r *wl.Region) {
	// TODO
}
