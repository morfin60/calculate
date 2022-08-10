package factory

import (
	"github.com/morfin60/calculate/internal/handlers"
)

type HandlerFactory struct {
}

const (
	WordsCount string = "words-count"
)

func NewHandlerFactory() *HandlerFactory {
	return &HandlerFactory{}
}

// Create handler instance
func (hf *HandlerFactory) Create(name string) handlers.Handler {
	switch name {
	case WordsCount:
		return handlers.NewWordsCount()
	}

	return nil
}
