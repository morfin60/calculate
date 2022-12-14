package factory

import (
	"github.com/morfin60/calculate/internal/handlers"
)

type HandlerFactory struct {
}

const (
	WordsCount    string = "words-count"
	TopWords      string = "top-10-words"
	TopCharacters string = "top-10-chars"
)

func NewHandlerFactory() *HandlerFactory {
	return &HandlerFactory{}
}

// Create handler instance
func (hf *HandlerFactory) Create(name string) handlers.Handler {
	switch name {
	case WordsCount:
		return handlers.NewWordsCount()
	case TopWords:
		return handlers.NewTopWords()
	case TopCharacters:
		return handlers.NewTopCharacters()
	}

	return nil
}
