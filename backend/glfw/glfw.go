package glfw

import "deedles.dev/moto/backend"

var Backend glfwBackend

func init() {
	backend.Register("glfw", Backend)
}

type glfwBackend struct {
	// TODO
}
