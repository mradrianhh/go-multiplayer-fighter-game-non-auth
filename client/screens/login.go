package screens

import (
	"fmt"

	"github.com/mradrianhh/go-multiplayer-fighter-game/client/network"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
)

// Login ..
type Login struct {
	Identifier string
}

// Show presents the login screen to the user.
func (login Login) Show(state *vars.State) error {
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

	response, err := network.MessageServer(models.NewMessage(vars.AUTHENTICATION, username+"\n"+password, vars.NEUTRAL))
	if err != nil {
		return err
	}

	fmt.Println(response.ResponseCode)
	if response.ResponseCode == vars.ACCEPTED {
		*state = vars.MAIN
	} else {
		*state = vars.LOGIN
	}
	return nil
}
