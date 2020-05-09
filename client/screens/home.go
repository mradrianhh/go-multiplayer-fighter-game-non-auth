package screens

import (
	"fmt"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
)

// Home ..
type Home struct {
	Identifier string
}

// Show presents the home menu to the user.
func (home Home) Show(state *vars.State) error {
	fmt.Println("1 - Login | 2 - Register | 0 - Exit Game")

	var response int
	_, err := fmt.Scan(&response)
	if err != nil {
		return err
	}

	switch response {
	case 1:
		*state = vars.LOGIN
	case 2:
		*state = vars.REGISTER
	case 0:
		os.Exit(0)
	default:
		*state = vars.HOME
		fmt.Println("Sorry, I can't understand...")
	}
	return nil
}
