package models

import (
	"fmt"
	"strings"

	"github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/vars"
)

// Message represents the object being passed between server and client.
type Message struct {
	MessageType  vars.MessageType
	Message      string
	ResponseCode vars.ResponseCode
	Token        string
}

// NilMessage is an empty message object.
var NilMessage = Message{}

// NewMessage returns a new message-instance.
func NewMessage(messageType vars.MessageType, message string, responseCode vars.ResponseCode, token string) Message {
	return Message{MessageType: messageType, Message: message, ResponseCode: responseCode, Token: token}
}

// Present prints the message to the screen.
func (message Message) Present(index int) {
	fmt.Printf("Message %v\n", index)
	fmt.Printf("\tMessage-type: %s\n", string(message.MessageType))
	fmt.Printf("\tMessage: %s\n", strings.ReplaceAll(message.Message, "\n", " "))
	fmt.Printf("\tResponse-code: %s\n", string(message.ResponseCode))
	fmt.Printf("\tToken: %s\n\n", message.Token)
}
