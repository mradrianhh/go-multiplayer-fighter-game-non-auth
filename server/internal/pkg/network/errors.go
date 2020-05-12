package network

import "errors"

// Network errors.

// ErrUnknownMessageType is thrown when the message-type passed with the message doesn't match with any of the message-types registered with the service.
var ErrUnknownMessageType = errors.New("unknown message type")

// ErrWrongMessageForType is thrown when the service can't process the message according to it's type.
var ErrWrongMessageForType = errors.New("can't process message according to type")

// Authentication errors.

// ErrWrongUsername is thrown when the "Login"-request's username doesn't match with any of the registered users.
var ErrWrongUsername = errors.New("wrong username")

// ErrWrongPassword is thrown when the "Login"-request's password is incorrect.
var ErrWrongPassword = errors.New("wrong password")

// ErrPlayerAlreadyLoggedIn is thrown when a player tries to connect, but it already logged in.
var ErrPlayerAlreadyLoggedIn = errors.New("player already logged in")

// ErrUnauthenticatedToken is thrown when an authenticated message is sent with an unauthenticated token.
var ErrUnauthenticatedToken = errors.New("token is not authenticated")

// Registration errors.

// ErrPlayerAlreadyExists is thrown when the "Registration"-request's username already exists in the userbase.
var ErrPlayerAlreadyExists = errors.New("player already exists")

// Events errors.

// ErrUnrecognizedEvent is thrown when the handler can't recognize the event sent.
var ErrUnrecognizedEvent = errors.New("event is not recognized")
