package screens

import (
	"fmt"

	cmodels "github.com/mradrianhh/go-multiplayer-fighter-game/client/internal/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/client/internal/pkg/network"
	"github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/vars"
)

var login = Login{Identifier: "LOGIN"}

// Login ..
type Login struct {
	Identifier string
}

// Show presents the login screen to the user.
func (login Login) Show(state *models.State) error {
	username := ""
	password := ""
	fmt.Print("Enter username: ")
	if _, err := fmt.Scan(&username); err != nil {
		return err
	}
	fmt.Print("Enter password: ")
	if _, err := fmt.Scan(&password); err != nil {
		return err
	}

	response, err := network.MessageServer(cmodels.NewMessage(vars.Authentication, username+"\n"+password))
	if err != nil {
		return err
	}

	if response.ResponseCode == vars.Accepted {
		state.Token = response.Token
		state.State = vars.Main
	} else {
		state.State = vars.Login
	}
	return nil
}

// GetLogin returns the login screen.
func GetLogin() *Login {
	return &login
}
