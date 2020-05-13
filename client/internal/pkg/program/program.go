package program

import (
	"fmt"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/client/internal/pkg/screens"
	"github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/vars"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/navigation"
)

var state models.State
var navigator navigation.Navigator

func init() {
	state.State = vars.Home
	navigator = navigation.NewNavigator()

	mainmenu := screens.GetMainMenu()
	navigator.AddScreen(mainmenu.Identifier, mainmenu)

	login := screens.GetLogin()
	navigator.AddScreen(login.Identifier, login)

	register := screens.GetRegister()
	navigator.AddScreen(register.Identifier, register)

	home := screens.GetHome()
	navigator.AddScreen(home.Identifier, home)
}

// Run initializes the game.
func Run() {
	Loop()
}

// Loop is the gameloop.
func Loop() {
	for {
		switch state.State {
		case vars.Main:
			err := navigator.RemoveScreensAndPush(&state, "MAINMENU")
			checkError(err)
		case vars.Login:
			err := navigator.RemoveScreensAndPush(&state, "LOGIN")
			checkError(err)
		case vars.Register:
			err := navigator.RemoveScreensAndPush(&state, "REGISTER")
			checkError(err)
		case vars.Home:
			err := navigator.RemoveScreensAndPush(&state, "HOME")
			checkError(err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
