package handlers

import (
	"sort"
	"strconv"
	"sync"

	"github.com/morfin60/calculate/internal/helpers"
)

type TopWords struct {
	BaseHandler
	wordCount map[string]int
}

func NewTopWords() *TopWords {
	return &TopWords{BaseHandler{make(chan string)}, map[string]int{}}
}

// Process incoming data
func (tw *TopWords) Process(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case word, ok := <-tw.dataChannel:
			if ok {
				tw.wordCount[word]++
			} else {
				return
			}
		}
	}
}

// Format result
func (tw *TopWords) Result() string {
	keys := make([]string, 0, len(tw.wordCount))

	for key := range tw.wordCount {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return tw.wordCount[keys[i]] > tw.wordCount[keys[j]]
	})

	table := helpers.NewTable(3)

	table.AddHeader([]string{"Rank", "Word", "Frequency"})

	//For each word from top 10 add row into table
	for i, key := range keys {
		rank := strconv.Itoa(i + 1)
		frequency := strconv.Itoa(tw.wordCount[key])
		table.AddRow([]string{rank, key, frequency})

		if i == 10 {
			break
		}
	}

	return table.ToString()
}
