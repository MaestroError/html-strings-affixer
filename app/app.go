package app

import (
	"fmt"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

func Bootstrap() {
	// configs and prepare app for work
}

func Start() {
	// Application runtime code goes here
}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}
