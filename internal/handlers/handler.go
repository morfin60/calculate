package handlers

import "sync"

type Handler interface {
	Process(wg *sync.WaitGroup)
	Data() chan<- string
	Result() string
}
