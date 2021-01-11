// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package thing

import (
	"github.com/beeceej/wire-sandbox/pkg/logger"
)

// Injectors from wire.go:

func InjectThing() Thing {
	stdoutLogger := logger.NewStdoutLogger()
	thing := NewThing(stdoutLogger)
	return thing
}

func InjectStderrThing() Thing {
	stderrLogger := logger.NewStderrLogger()
	thing := NewThing(stderrLogger)
	return thing
}
