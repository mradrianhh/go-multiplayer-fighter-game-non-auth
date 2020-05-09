package network

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/models"
)

const (
	service = "0.0.0.0:1200"
	network = "tcp"
)

// MessageServer sends a message to the server and returns the response.
func MessageServer(message models.Message) (models.Message, error) {
	conn, err := net.Dial("tcp", service)
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

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
