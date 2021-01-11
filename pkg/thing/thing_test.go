// Use `go generate` generate the mocks needed for this test suite
//go:generate mockgen -source=../logger/logger.go -destination=../mocks/logger.go -package mocks

package thing

import (
	"testing"

	"github.com/beeceej/wire-sandbox/pkg/mocks"
	"github.com/golang/mock/gomock"
)

func TestThing(t *testing.T) {
	// SETUP
	c := gomock.NewController(t)
	mockLogger := mocks.NewMockLogger(c)
	thing := NewThing(mockLogger)

	// SYSTEM UNDER TEST
	defer func() {
		thing.Say("Hello, World")
	}()

	// ASSERT
	mockLogger.EXPECT().Log("Hello, World")
}
