package models

// MessageType is a string representation of the type of message. I.E "AUTHENTICATION", "CONFIRMATION", etc.
type MessageType string

// Message types.
const (
	CONFIRMATION   = "CONFIRMATION"
	AUTHENTICATION = "AUTHENTICATION"
	EVENT          = "EVENT"
)

// ResponseCode ...
type ResponseCode string

// Response codes
const (
	ACCEPTED    = "ACCEPTED"
	NOTACCEPTED = "NOT ACCEPTED"
	NEUTRAL     = "NEUTRAL"
)

// Message represents the object being passed between server and client.
type Message struct {
	MessageType  MessageType
	Message      string
	ResponseCode ResponseCode
}

// NewMessage returns a new message-instance.
func NewMessage(messageType MessageType, message string, responseCode ResponseCode) Message {
	return Message{MessageType: messageType, Message: message, ResponseCode: responseCode}
}

// ErrorMessage is the default message when something goes wrong.
var ErrorMessage = Message{MessageType: CONFIRMATION, Message: "error", ResponseCode: NOTACCEPTED}
