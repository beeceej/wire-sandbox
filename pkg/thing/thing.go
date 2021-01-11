package thing

import (
	"github.com/beeceej/wire-sandbox/pkg/logger"
)

type (
	// Thing is a struct that can say stuff with the given logger
	Thing struct {
		logger logger.Logger
	}
)

// Say uses the provided logger to write the given message
func (t *Thing) Say(msg string) {
	t.logger.Log(msg)
}

// NewThing is a concrete provider of a Thing
func NewThing(l logger.Logger) Thing {
	return Thing{logger: l}
}
