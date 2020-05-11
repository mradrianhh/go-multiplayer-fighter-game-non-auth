package network

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
	"github.com/mradrianhh/go-multiplayer-fighter-game/server/pkg/models"
)

var playerDatabase map[string]models.Player

// Players holds all the players connected to the session.
var Players map[string]models.Player

// Messages holds a record of all the messages passed.
var Messages []models.Message

// Log holds a record of all logs written.
var Log []string

// ErrorLog holds a record of all errors encountered.
var ErrorLog []error

// Conns holds a record of all the connections.
var Conns map[string]net.Conn

func init() {
	playerDatabase = make(map[string]models.Player)
	Players = make(map[string]models.Player)
	Conns = make(map[string]net.Conn)

	playerDatabase["username"] = models.NewPlayer("username", "password")
}

// Listen opens a connection that listens for messages from a client.
func Listen() {
	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			addError(err)
			continue // if there is an error creating a connection, we'll move on and let the client deal with it. They can retry.
		}

		decoder := gob.NewDecoder(conn)
		var message models.Message
		decoder.Decode(&message) // Read into the "message"-variable from the stream.
		Messages = append(Messages, message)
		err = handleMessage(message, conn) // handle the message.
		if err != nil {
			addError(err)
			writeLog(err.Error())
		}
	}
}

func handleMessage(message models.Message, conn net.Conn) error {
	switch message.MessageType {
	case vars.Authentication:
		return handleAuthentication(message, conn)
	case vars.Registration:
		return handleRegistration(message, conn)
	case vars.Event:
		return handleEvent(message, conn)
	default:
		return ErrUnknownMessageType
	}
}

func handleAuthentication(message models.Message, conn net.Conn) error {
	strs := strings.Split(message.Message, "\n")
	if len(strs) != 2 {
		return ErrWrongMessageForType
	}

	providedUsername := strs[0]
	providedPassword := strs[1]
	message, err := authenticate(providedUsername, providedPassword)
	if err == nil {
		Conns[providedUsername] = conn // if the  user gets authenticated, it means he's logged in with "providedUsername", in which case we add his conn to the list.
		sendMessage(message, conn)
		return nil
	}
	return err
}

func authenticate(providedUsername, providedPassword string) (models.Message, error) {
	if _, exist := Players[providedUsername]; exist {
		return msgPlayerAlreadyLoggedIn, ErrPlayerAlreadyLoggedIn
	}

	player, exist := playerDatabase[providedUsername]
	if exist == false {
		return msgWrongUsername, ErrWrongUsername
	} else if player.Password != providedPassword {
		return msgWrongPassword, ErrWrongPassword
	} else {
		Players[player.Username] = player
		return models.NewMessage(vars.Confirmation, player.Username+" authenticated.", vars.Accepted, player.Username), nil
	}
}

func handleRegistration(message models.Message, conn net.Conn) error {
	strs := strings.Split(message.Message, "\n")
	if len(strs) != 2 {
		return ErrWrongMessageForType
	}

	username := strs[0]
	password := strs[1]

	message, err := registrate(username, password)
	sendMessage(message, conn)
	return err
}

func registrate(username, password string) (models.Message, error) {
	if _, exist := playerDatabase[username]; exist == false {
		playerDatabase[username] = models.NewPlayer(username, password)

		return newMsgPlayerRegistered(username), nil
	}

	return msgPlayerAlreadyExists, ErrPlayerAlreadyExists
}

func handleEvent(message models.Message, conn net.Conn) error {
	_, exist := Players[message.Token]
	if exist == false {
		sendMessage(newMsgUnauthenticatedToken(message.Token), conn)
		return ErrUnauthenticatedToken
	}

	switch message.Message {
	case vars.LoggedOut:
		delete(Players, message.Token)
		sendMessage(msgAccepted, conn)
		return nil
	default:
		return ErrUnrecognizedEvent
	}
}

func sendMessage(message models.Message, conn net.Conn) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(message)
	Messages = append(Messages, message)
}

func checkError(err error) {
	if err != nil {
		addError(err)
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func writeLog(text string) {
	Log = append(Log, text)
}

func addError(err error) {
	ErrorLog = append(ErrorLog, err)
}
