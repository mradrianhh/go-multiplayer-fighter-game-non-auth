package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/models"
)

var users map[string]string

func init() {
	users = make(map[string]string)
	users["adrianhh"] = "fjellhomse123"
}

func main() {
	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		decoder := gob.NewDecoder(conn)
		var message models.Message
		decoder.Decode(&message)
		err = handleMessage(message, conn)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func handleMessage(message models.Message, conn net.Conn) error {
	switch message.MessageType {
	case models.AUTHENTICATION:
		strs := strings.Split(message.Message, "\n")
		if len(strs) != 2 {
			return errors.New("can't process message according to type. 'Authentication'-message should only contain 2 strings")
		}
		providedUsername := strs[0]
		providedPassword := strs[1]
		response, err := authenticate(providedUsername, providedPassword)
		sendMessage(models.NewMessage(models.CONFIRMATION, string(response), models.ACCEPTED), conn)
		return err
	case models.EVENT:
		return nil
	default:
		return errors.New("unknown message type")
	}
}

func authenticate(providedUsername, providedPassword string) (models.ResponseCode, error) {
	password := users[providedUsername]
	if password == "" {
		return models.NOTACCEPTED, errors.New("wrong username")
	} else if password != providedPassword {
		return models.NOTACCEPTED, errors.New("wrong password")
	} else {
		return models.ACCEPTED, nil
	}
}

func sendMessage(message models.Message, conn net.Conn) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(message)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
