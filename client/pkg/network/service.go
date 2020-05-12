package network

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/client/pkg/models"
)

const (
	service = "0.0.0.0:1200"
	network = "tcp"
)

// Run starts the service.
func Run() {
	go listen()
}

// MessageServer sends a message to the server and returns the response.
func MessageServer(message models.Message) (models.Message, error) {
	conn, err := net.Dial(network, service)
	if err != nil {
		return models.ErrorMessage, err
	}

	encoder := gob.NewEncoder(conn)
	encoder.Encode(message)

	var response models.Message
	decoder := gob.NewDecoder(conn)
	decoder.Decode(&response)
	return response, nil
}

func listen() {
	tcpAddr, err := net.ResolveTCPAddr(network, service)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP(network, tcpAddr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		decoder := gob.NewDecoder(conn)
		var message models.Message
		decoder.Decode(&message)
		fmt.Println(message.MessageType)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
