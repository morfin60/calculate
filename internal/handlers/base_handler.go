package handlers

type BaseHandler struct {
	dataChannel   chan string
	resultChannel chan string
}

// Get data channel
func (bh *BaseHandler) Data() chan<- string {
	return bh.dataChannel
}

// Get result channel
func (bh *BaseHandler) Result() <-chan string {
	return bh.resultChannel
}
