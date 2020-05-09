package vars

// State represents the current state of the game.
type State string

// Game states
const (
	HOME     = State("HOME")
	LOGIN    = State("LOGIN")
	MAIN     = State("MAIN")
	REGISTER = State("REGISTER")
)

// MessageType is a string representation of the type of message. I.E "AUTHENTICATION", "CONFIRMATION", etc.
type MessageType string

// Message types.
const (
	CONFIRMATION   = "CONFIRMATION"
	AUTHENTICATION = "AUTHENTICATION"
	EVENT          = "EVENT"
	REGISTRATION   = "REGISTRATION"
)

// ResponseCode ...
type ResponseCode string

// Response codes
const (
	ACCEPTED    = "ACCEPTED"
	NOTACCEPTED = "NOT ACCEPTED"
	NEUTRAL     = "NEUTRAL"
	SUCCEEDED   = "SUCCEEDED"
	FAILED      = "FAILED"
)
