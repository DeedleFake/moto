package moto

import "deedles.dev/moto/backend"

type Renderer struct {
	impl     backend.Renderer
	listener rendererListener

	rendering, renderingWithBuffer bool
}

func AutocreateRenderer(backend *Backend) (*Renderer, error) {
	// TODO
	return nil, nil
}

func (r *Renderer) InitDisplay(display *Display) {
	// TODO
}

type rendererListener interface {
	Destroy()
}
