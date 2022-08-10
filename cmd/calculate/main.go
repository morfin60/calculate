package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/morfin60/calculate/internal/handlers"
	"github.com/morfin60/calculate/internal/handlers/factory"
	"github.com/morfin60/calculate/internal/helpers"
)

func main() {
	handlersFactory := factory.NewHandlerFactory()
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	handlers := map[string]handlers.Handler{}

	scanner.Split(helpers.SplitText)

	// For each argument get handler if exists and process input data
	for _, name := range os.Args[1:] {
		if handler := handlersFactory.Create(name); handler != nil {
			handlers[name] = handler
			go handler.Process()
		}
	}

	// Scan input for words
	for scanner.Scan() {
		for name, _ := range handlers {
			handlers[name].Data() <- scanner.Text()
		}
	}

	// Close data channels when out of data
	for name := range handlers {
		close(handlers[name].Data())
	}

	// Read results from channels
	for name, _ := range handlers {
		fmt.Print(<-handlers[name].Result())
	}

}
