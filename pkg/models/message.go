package models

import "github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"

// Message represents the object being passed between server and client.
type Message struct {
	MessageType  vars.MessageType
	Message      string
	ResponseCode vars.ResponseCode
}

// NewMessage returns a new message-instance.
func NewMessage(messageType vars.MessageType, message string, responseCode vars.ResponseCode) Message {
	return Message{MessageType: messageType, Message: message, ResponseCode: responseCode}
}

// ErrorMessage is the default message when something goes wrong.
var ErrorMessage = Message{MessageType: vars.CONFIRMATION, Message: "error", ResponseCode: vars.NOTACCEPTED}
