package handlers

import (
	"sort"
	"strconv"
	"sync"

	"github.com/morfin60/calculate/internal/helpers"
)

type wordInfo struct {
	count int
	order int
}

type TopWords struct {
	BaseHandler
	wordCount map[string]*wordInfo
}

func NewTopWords() *TopWords {
	return &TopWords{BaseHandler{make(chan string)}, map[string]*wordInfo{}}
}

// Process incoming data
func (tw *TopWords) Process(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case word, ok := <-tw.dataChannel:
			if ok {
				if _, ok := tw.wordCount[word]; !ok {
					tw.wordCount[word] = &wordInfo{count: 0, order: len(tw.wordCount)}
				}
				tw.wordCount[word].count += 1
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

	sort.SliceStable(keys, func(i, j int) bool {
		if tw.wordCount[keys[i]].count == tw.wordCount[keys[j]].count {
			return tw.wordCount[keys[i]].order > tw.wordCount[keys[j]].order
		}

		return tw.wordCount[keys[i]].count > tw.wordCount[keys[j]].count
	})

	table := helpers.NewTable(3)

	table.AddHeader([]string{"Rank", "Word", "Frequency"})

	//For each word from top 10 add row into table
	for i, key := range keys {
		rank := strconv.Itoa(i + 1)
		frequency := strconv.Itoa(tw.wordCount[key].count)
		table.AddRow([]string{rank, key, frequency})

		if i+1 == 10 {
			break
		}
	}

	return table.ToString()
}
