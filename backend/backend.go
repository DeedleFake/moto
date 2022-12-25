package backend

import (
	"golang.org/x/exp/maps"
)

var backends map[string]Backend

func Register(name string, backend Backend) {
	backends[name] = backend
}

func Names() []string {
	return maps.Keys(backends)
}

type Backend interface {
	//Start() error
	//Destroy()
	//Session() *Session
	//PresentationClock() *Clock
	//DRM() *os.File
	//BufferCaps() uint32
}

type Renderer interface {
	//BindBuffer(*Buffer)
	//Begin(width, height uint32)
	//End()
	//Clear(color.Color)
	//Scissor(image.Rectangle)
	//RenderSubtextureWithMatrix(*Texture, image.Rectangle, f32.Mat3, float32) error
	//RenderQuadWithMatrix(color.Color, f32.Mat3)
	//ShmTextureFormats() []uint32
	//DmabufTextureFormats() DrmFormatSet
	//RenderFormats() DrmFormatSet
	//PreferredReadFormat() uint32
	//ReadPixels(data []byte, fmt, flags, stride, width, height uint32, src, dst image.Point) error
	//Destroy()
	//DRM() *os.File
	//RenderBufferCaps() uint32
	//TextureFromBuffer(*Buffer) *Texture
}
