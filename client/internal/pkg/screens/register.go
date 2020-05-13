package screens

import (
	"fmt"

	cmodels "github.com/mradrianhh/go-multiplayer-fighter-game/client/internal/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/client/internal/pkg/network"
	"github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/vars"
)

var register = Register{Identifier: "REGISTER"}

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

	response, err := network.MessageServer(cmodels.NewMessage(vars.Registration, username+"\n"+password))
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

// GetRegister returns the register screen.
func GetRegister() *Register {
	return &register
}
