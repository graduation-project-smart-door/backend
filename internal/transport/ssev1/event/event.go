package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	"smart-door/internal/domain"

	"github.com/gorilla/mux"
)

type Broker struct {
	Notifier chan []byte

	newClients chan chan []byte

	closingClients chan chan []byte

	clients map[chan []byte]bool
}

func NewBroker() *Broker {
	broker := &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}

	go broker.listen()

	return broker
}

func (broker *Broker) Register(router *mux.Router) {
	router.HandleFunc("", broker.ServeHTTP)
}

func (broker *Broker) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	flusher, ok := writer.(http.Flusher)

	if !ok {
		return
	}

	messageChan := make(chan []byte)
	broker.newClients <- messageChan
	defer func() {
		broker.closingClients <- messageChan
	}()

	writer.Header().Set("Content-Type", "text/event-stream")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Connection", "keep-alive")

	go func() {
		<-request.Context().Done()
		broker.closingClients <- messageChan
	}()

	for {
		select {
		case message := <-messageChan:
			fmt.Fprintf(writer, "data: %s\n\n", message)
			flusher.Flush()
		case <-request.Context().Done():
			broker.closingClients <- messageChan
			return
		}
	}
}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:
			broker.clients[s] = true
		case s := <-broker.closingClients:
			delete(broker.clients, s)
		case event := <-broker.Notifier:
			for clientMessageChan, _ := range broker.clients {
				clientMessageChan <- event
			}
		}
	}
}

func (broker *Broker) ToMessage(message domain.Event) error {
	body, errMarshal := json.Marshal(message)
	if errMarshal != nil {
		return errMarshal
	}

	broker.Notifier <- body
	return nil
}
