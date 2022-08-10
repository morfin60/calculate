package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/morfin60/calculate/internal/handlers"
	"github.com/morfin60/calculate/internal/handlers/factory"
	"github.com/morfin60/calculate/internal/helpers"
)

func main() {
	// If no handlers specified just show usage and exit
	if len(os.Args) == 1 {
		println("Usage: calculate [handlers]")

		return
	}

	wg := sync.WaitGroup{}
	handlersFactory := factory.NewHandlerFactory()
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	handlers := []handlers.Handler{}

	scanner.Split(helpers.SplitText)

	// For each argument get handler if exists and process input data
	for _, name := range os.Args[1:] {
		if handler := handlersFactory.Create(name); handler != nil {
			wg.Add(1)
			handlers = append(handlers, handler)
			go handler.Process(&wg)
		} else {
			println("Handler " + name + " does not exist")
		}
	}

	// No handlers created just exit
	if len(handlers) == 0 {
		return
	}

	// Scan input for words
	for scanner.Scan() {
		for key := range handlers {
			handlers[key].Data() <- scanner.Text()
		}
	}

	// Close data channels when out of data
	for key := range handlers {
		close(handlers[key].Data())
	}

	wg.Wait()

	// Read results from channels
	for name := range handlers {
		fmt.Printf("%s\n\n", handlers[name].Result())
	}
}
