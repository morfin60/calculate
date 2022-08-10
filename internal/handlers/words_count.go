package handlers

import (
	"fmt"
	"sync"
)

type WordsCount struct {
	BaseHandler
	wordsCount int
}

func NewWordsCount() *WordsCount {
	return &WordsCount{BaseHandler{make(chan string)}, 0}
}

// Process incoming data
func (wc *WordsCount) Process(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case _, ok := <-wc.dataChannel:
			if ok {
				wc.wordsCount++
			} else {
				return
			}
		}
	}

}

// Format result
func (wc *WordsCount) Result() string {
	return fmt.Sprintf("WordsCount:\n     %d\n", wc.wordsCount)
}
