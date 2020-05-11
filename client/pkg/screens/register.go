package screens

import (
	"fmt"

	"github.com/mradrianhh/go-multiplayer-fighter-game/client/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/client/pkg/network"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
)

// Register ..
type Register struct {
	Identifier string
}

// Show presents the register screen to the user.
func (register Register) Show(state *models.State) error {
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

	response, err := network.MessageServer(models.NewMessage(vars.Registration, username+"\n"+password))
	if err != nil {
		return err
	}

	if response.ResponseCode == vars.Accepted {
		state.State = vars.Home
	} else {
		state.State = vars.Register
	}
	return nil
}
