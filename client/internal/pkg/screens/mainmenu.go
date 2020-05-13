package screens

import (
	"fmt"

	cmodels "github.com/mradrianhh/go-multiplayer-fighter-game/client/internal/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/client/internal/pkg/network"
	"github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/vars"
)

var mainmenu = MainMenu{Identifier: "MAINMENU"}

// MainMenu represents the main menu.
type MainMenu struct {
	Identifier string
}

// Show presents the main menu to the user.
func (mainMenu MainMenu) Show(state *models.State) error {
	fmt.Println("1 - Start Game | 2 - Logout")

	var response int
	_, err := fmt.Scan(&response)
	if err != nil {
		return err
	}

	switch response {
	case 1:
		fmt.Println("Starting game...")
		state.State = vars.Home
	case 2:
		response, err := network.MessageServer(cmodels.NewAuthenticatedMessage(vars.Event, vars.LoggedOut, state.Token))
		if err != nil {
			state.State = vars.Main
			fmt.Println("Can't log out, error encountered. Try again.")
		}
		if response.ResponseCode != vars.Accepted {
			state.State = vars.Main
			fmt.Printf("Can't log out. Response-code: %s. Try again.\n", string(response.ResponseCode))
		}
		fmt.Println("Logging out...")
		state.State = vars.Home
	default:
		state.State = vars.Main
		fmt.Println("Sorry, I can't understand...")
	}
	return nil
}

// GetMainMenu returns the main-menu screen.
func GetMainMenu() *MainMenu {
	return &mainmenu
}
