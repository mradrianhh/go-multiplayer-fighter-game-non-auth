package models

import (
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
)

// Message represents the object being passed between server and client.
type Message struct {
	MessageType  vars.MessageType
	Message      string
	ResponseCode vars.ResponseCode
	Token        string
}

// NewMessage returns a new message-instance.
func NewMessage(messageType vars.MessageType, message string) Message {
	return Message{MessageType: messageType, Message: message}
}

// NewAuthenticatedMessage returns a new message-instance with a token.
func NewAuthenticatedMessage(messageType vars.MessageType, message string, token string) Message {
	return Message{MessageType: messageType, Message: message, Token: token}
}

// ErrorMessage is the default message when something goes wrong.
var ErrorMessage = Message{MessageType: vars.Confirmation, Message: "error", ResponseCode: vars.NotAccepted}
