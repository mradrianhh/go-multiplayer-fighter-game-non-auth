package screens

import (
	"fmt"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/client/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
)

// Home ..
type Home struct {
	Identifier string
}

// Show presents the home menu to the user.
func (home Home) Show(state *models.State) error {
	fmt.Println("1 - Login | 2 - Register | 0 - Exit Game")

	var response int
	_, err := fmt.Scan(&response)
	if err != nil {
		return err
	}

	switch response {
	case 1:
		state.State = vars.Login
	case 2:
		state.State = vars.Register
	case 0:
		os.Exit(0)
	default:
		state.State = vars.Home
		fmt.Println("Sorry, I can't understand...")
	}
	return nil
}