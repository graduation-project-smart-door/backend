package event

import (
	"fmt"
	"net/http"

	"smart-door/internal/apperror"

	"github.com/gorilla/mux"
)

var messageChan chan string

type Handler struct {
	policy EventPolicy
}

func NewHandler(policy EventPolicy) *Handler {
	return &Handler{policy: policy}
}

func (handler *Handler) Register(router *mux.Router) {
	router.HandleFunc("", apperror.Middleware(handler.newEvents))
}

func (handler *Handler) newEvents(writer http.ResponseWriter, request *http.Request) error {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "text/event-stream")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Connection", "keep-alive")

	messageChan := make(chan string)

	defer func() {
		close(messageChan)
		messageChan = nil
	}()

	flusher, ok := writer.(http.Flusher)
	if !ok {
		fmt.Println("error init")
		return nil
	}

	for {
		select {
		case message := <-messageChan:
			fmt.Fprintf(writer, "data: %s\n\n", message)
			flusher.Flush()
		case <-request.Context().Done():
			fmt.Println("client closed")
			return nil
		}
	}
}
