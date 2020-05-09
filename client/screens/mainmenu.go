package screens

import (
	"fmt"

	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
)

// MainMenu represents the main menu.
type MainMenu struct {
	Identifier string
}

// Show presents the main menu to the user.
func (mainMenu MainMenu) Show(state *vars.State) error {
	fmt.Println("1 - Start Game | 2 - Logout")

	var response int
	_, err := fmt.Scan(&response)
	if err != nil {
		return err
	}

	switch response {
	case 1:
		fmt.Println("Starting game...")
		*state = vars.HOME
	case 2:
		*state = vars.HOME
	default:
		*state = vars.MAIN
		fmt.Println("Sorry, I can't understand...")
	}
	return nil
}
