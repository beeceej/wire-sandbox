// +build wireinject

package thing

import (
	"github.com/beeceej/wire-sandbox/pkg/logger"
	"github.com/google/wire"
)

// InjectThing is a function that tells wire how to build a thing.Thing
func InjectThing() Thing {
	panic(wire.Build(NewThing, logger.Stdoutset))
}

// InjectStderrThing is a function that tells wire how to build a thing.Thing
func InjectStderrThing() Thing {
	panic(wire.Build(NewThing, logger.Stderrset))
}
