package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"deedles.dev/moto"
)

type server struct {
	display  *moto.Display
	backend  *moto.Backend
	renderer *moto.Renderer
}

func (s *server) init() {
	display, err := moto.NewDisplay()
	if err != nil {
		log.Fatalf("new display: %v", err)
	}
	s.display = display

	backend, err := moto.AutocreateBackend(s.display)
	if err != nil {
		log.Fatalf("create backend: %v", err)
	}
	s.backend = backend

	renderer, err := moto.AutocreateRenderer(s.backend)
	if err != nil {
		log.Fatalf("create renderer: %v", err)
	}
	s.renderer = renderer
	s.renderer.InitDisplay(s.display)

	moto.CreateCompositor(s.display, s.renderer)
	//moto.NewDataDeviceManager(s.display)

	//s.outputLayout = s.CreateOutputLayout()
	//s.backend.NewOutput = s.onNewOutput

	//s.scene = moto.CreateScene()
	//s.scene.AttachOutputLayout(s.outputLayout)
}

func (s *server) run(ctx context.Context) {
	s.display.Run(ctx)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var s server
	s.init()
	s.run(ctx)
}
