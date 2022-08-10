package handlers

type Handler interface {
	Process()
	Data() chan<- string
	Result() <-chan string
}
