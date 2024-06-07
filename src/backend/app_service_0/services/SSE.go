package services

import (
	"fmt"
	"net/http"
)

var messageChan chan string

func ServerSentEvent(msg string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	messageChan <- msg

	flusher := w.(http.Flusher)
	messageChan = make(chan string)

	defer func() {
		close(messageChan)
		messageChan = nil
	}()

	for {
		select {

		case message := <-messageChan:
			fmt.Fprintf(w, "data: %s\n\n", message)
			flusher.Flush()

		case <-r.Context().Done():
			return
		}
	}
}
