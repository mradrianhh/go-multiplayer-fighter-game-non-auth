package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
)

var users map[string]string

func init() {
	users = make(map[string]string)
	users["username"] = "password"
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
	case vars.AUTHENTICATION:
		strs := strings.Split(message.Message, "\n")
		if len(strs) != 2 {
			return errors.New("can't process message according to type. 'Authentication'-message should only contain 2 strings")
		}

		providedUsername := strs[0]
		providedPassword := strs[1]
		responseCode, err := authenticate(providedUsername, providedPassword)
		sendMessage(models.NewMessage(vars.CONFIRMATION, string(responseCode), responseCode), conn)
		return err
	case vars.REGISTRATION:
		strs := strings.Split(message.Message, "\n")
		if len(strs) != 2 {
			return errors.New("can't process message according to type. 'Registration'-message should only contain 2 strings")
		}

		username := strs[0]
		password := strs[1]

		responseCode, err := registrate(username, password)
		sendMessage(models.NewMessage(vars.CONFIRMATION, string(responseCode), responseCode), conn)
		return err
	case vars.EVENT:
		return nil
	default:
		return errors.New("unknown message type")
	}
}

func authenticate(providedUsername, providedPassword string) (vars.ResponseCode, error) {
	password := users[providedUsername]
	if password == "" {
		return vars.NOTACCEPTED, errors.New("wrong username")
	} else if password != providedPassword {
		return vars.NOTACCEPTED, errors.New("wrong password")
	} else {
		return vars.ACCEPTED, nil
	}
}

func registrate(username, password string) (vars.ResponseCode, error) {
	users[username] = password

	return vars.SUCCEEDED, nil
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
