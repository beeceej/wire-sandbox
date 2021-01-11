package logger

import (
	"io"
	"os"

	"github.com/google/wire"
)

type (
	// Logger defines the behavior of Logging
	Logger interface {
		Log(string)
	}

	// StdoutLogger logs to stdout
	StdoutLogger struct {
		io.Writer
	}

	// StderrLogger logs to stderr
	StderrLogger struct {
		io.Writer
	}
)

func log(w io.Writer, msg []byte) {
	w.Write(msg)
}

// Log logs a message to the provider writer
func (i *StdoutLogger) Log(msg string) {
	log(i, []byte(msg))
}

// Log logs a message to the provider writer
func (i *StderrLogger) Log(msg string) {
	log(i, []byte(msg))
}

// NewStdoutLogger is a Provider for a *StdoutLogger
func NewStdoutLogger() *StdoutLogger {
	return &StdoutLogger{os.Stdout}
}

// NewStderrLogger is a Provider for a *StderrLogger
func NewStderrLogger() *StderrLogger {
	return &StderrLogger{os.Stderr}
}

// Wire Sets for
var (
	Stdoutset = wire.NewSet(wire.Bind(new(Logger), new(*StdoutLogger)), NewStdoutLogger)
	Stderrset = wire.NewSet(wire.Bind(new(Logger), new(*StderrLogger)), NewStderrLogger)
)
