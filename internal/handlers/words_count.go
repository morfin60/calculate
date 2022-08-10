package handlers

import (
	"fmt"
)

type WordsCount struct {
	BaseHandler
	wordsCount int
}

func NewWordsCount() *WordsCount {
	return &WordsCount{BaseHandler{make(chan string), make(chan string)}, 0}
}

// Process incoming data
func (wch *WordsCount) Process() {
	for {
		select {
		case _, ok := <-wch.dataChannel:
			if ok {
				wch.wordsCount++
			} else {
				wch.resultChannel <- wch.format()

				return
			}
		}
	}

}

// Format result
func (wch *WordsCount) format() string {
	return fmt.Sprintf("WordsCount:\n     %d\n", wch.wordsCount)
}
