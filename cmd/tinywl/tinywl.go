package main

import (
	"context"
	"os"
	"os/signal"

	"deedles.dev/moto"
	"golang.org/x/exp/slog"
)

type server struct {
	display *moto.Display
}

func (s *server) init() {
	display, err := moto.NewDisplay()
	if err != nil {
		slog.Error("new display", err)
		os.Exit(1)
	}
	s.display = display

	//backend, err := moto.AutocreateBackend(s.display)
	//if err != nil {
	//	slog.Error("create backend", err)
	//	os.Exit(1)
	//}
	//s.backend = backend

	//renderer, err := moto.AutocreateRenderer(s.backend)
	//if err != nil {
	//	slog.Error("create renderer", err)
	//	os.Exit(1)
	//}
	//s.renderer = renderer
	//s.renderer.InitDisplay(s.display)

	//moto.NewCompositor(s.display, s.renderer)
	//moto.NewDataDeviceManager(s.display)

	//s.outputLayout = moto.NewOutputLayout()

	//s.backend.NewOutput = s.onNewOutput

	//s.scene = moto.NewScene()
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
