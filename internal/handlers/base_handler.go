package handlers

type BaseHandler struct {
	dataChannel chan string
}

// Get data channel
func (bh *BaseHandler) Data() chan<- string {
	return bh.dataChannel
}
