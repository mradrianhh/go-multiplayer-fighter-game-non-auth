package network

import (
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
	"github.com/mradrianhh/go-multiplayer-fighter-game/server/pkg/models"
)

// Authentication messages.
var msgWrongUsername = models.NewMessage(vars.Confirmation, ErrWrongUsername.Error(), vars.NotAccepted, "")
var msgWrongPassword = models.NewMessage(vars.Confirmation, ErrWrongPassword.Error(), vars.NotAccepted, "")
var msgPlayerAlreadyLoggedIn = models.NewMessage(vars.Confirmation, ErrPlayerAlreadyLoggedIn.Error(), vars.NotAccepted, "")

// Registration messages.
var msgAccepted = models.NewMessage(vars.Confirmation, "", vars.Accepted, "")
var msgPlayerAlreadyExists = models.NewMessage(vars.Confirmation, ErrPlayerAlreadyExists.Error(), vars.NotAccepted, "")

func newMsgPlayerRegistered(username string) models.Message {
	return models.NewMessage(vars.Confirmation, username+" registered.", vars.Accepted, "")
}

// Event messages.
func newMsgUnauthenticatedToken(token string) models.Message {
	return models.NewMessage(vars.Confirmation, token+" is not registered as an authenticated token.", vars.NotAccepted, token)
}
