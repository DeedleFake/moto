package moto

import (
	"image"
	"image/color"
	"os"

	"golang.org/x/image/math/f32"
)

type Renderer struct {
	impl     RendererImpl
	listener rendererListener

	rendering, renderingWithBuffer bool
}

type RendererImpl interface {
	BindBuffer(*Buffer)
	Begin(width, height uint32)
	End()
	Clear(color.Color)
	Scissor(image.Rectangle)
	RenderSubtextureWithMatrix(*Texture, image.Rectangle, f32.Mat3, float32) error
	RenderQuadWithMatrix(color.Color, f32.Mat3)
	ShmTextureFormats() []uint32
	DmabufTextureFormats() DrmFormatSet
	RenderFormats() DrmFormatSet
	PreferredReadFormat() uint32
	ReadPixels(data []byte, fmt, flags, stride, width, height uint32, src, dst image.Point) error
	Destroy()
	DRM() *os.File
	RenderBufferCaps() uint32
	TextureFromBuffer(*Buffer) *Texture
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

type Buffer struct {
	// TODO
}

type Texture struct {
	// TODO
}

type DrmFormatSet struct {
	// TODO
}
