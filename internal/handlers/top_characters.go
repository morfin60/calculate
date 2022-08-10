package handlers

import (
	"sort"
	"strconv"
	"sync"

	"github.com/morfin60/calculate/internal/helpers"
)

type TopCharacters struct {
	BaseHandler
	charactersCount map[string]int
}

func NewTopCharacters() *TopCharacters {
	return &TopCharacters{BaseHandler{make(chan string)}, map[string]int{}}
}

// Process incoming data
func (tc *TopCharacters) Process(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case word, ok := <-tc.dataChannel:
			if ok {
				for _, char := range word {
					tc.charactersCount[string(char)]++
				}
			} else {
				return
			}
		}
	}
}

// Format result
func (tc *TopCharacters) Result() string {
	keys := make([]string, 0, len(tc.charactersCount))

	for key := range tc.charactersCount {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return tc.charactersCount[keys[i]] > tc.charactersCount[keys[j]]
	})

	table := helpers.NewTable(3)

	table.AddHeader([]string{"Rank", "Character", "Frequency"})

	//For each word from top 10 add row into table
	for i, key := range keys {
		rank := strconv.Itoa(i + 1)
		frequency := strconv.Itoa(tc.charactersCount[key])
		table.AddRow([]string{rank, key, frequency})

		if i+1 == 10 {
			break
		}
	}

	return table.ToString()
}
