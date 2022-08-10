package handlers

import (
	"sort"
	"strconv"
	"sync"

	"github.com/morfin60/calculate/internal/helpers"
)

type charInfo struct {
	count int
	order int
}

type TopCharacters struct {
	BaseHandler
	charactersCount map[string]*charInfo
}

func NewTopCharacters() *TopCharacters {
	return &TopCharacters{BaseHandler{make(chan string)}, map[string]*charInfo{}}
}

// Process incoming data
func (tc *TopCharacters) Process(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case word, ok := <-tc.dataChannel:
			if ok {
				for _, char := range word {
					key := string(char)
					if _, ok := tc.charactersCount[key]; !ok {
						tc.charactersCount[key] = &charInfo{count: 0, order: len(tc.charactersCount)}
					}
					tc.charactersCount[key].count += 1
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

	sort.SliceStable(keys, func(i, j int) bool {
		if tc.charactersCount[keys[i]].count == tc.charactersCount[keys[j]].count {
			return tc.charactersCount[keys[i]].order == tc.charactersCount[keys[j]].order
		}

		return tc.charactersCount[keys[i]].count > tc.charactersCount[keys[j]].count
	})

	table := helpers.NewTable(3)

	table.AddHeader([]string{"Rank", "Character", "Frequency"})

	//For each word from top 10 add row into table
	for i, key := range keys {
		rank := strconv.Itoa(i + 1)
		frequency := strconv.Itoa(tc.charactersCount[key].count)
		table.AddRow([]string{rank, key, frequency})

		if i+1 == 10 {
			break
		}
	}

	return table.ToString()
}
