package navigation

import (
	"github.com/mradrianhh/go-multiplayer-fighter-game/client/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/client/pkg/screens"
)

var navigationStack Stack

// Screens is a key-value collection of the screens registered.
var Screens map[string]Screen

func init() {
	mainmenu := screens.MainMenu{Identifier: "MAINMENU"}
	home := screens.Home{Identifier: "HOME"}
	login := screens.Login{Identifier: "LOGIN"}
	register := screens.Register{Identifier: "REGISTER"}

	Screens = make(map[string]Screen)
	Screens[mainmenu.Identifier] = mainmenu
	Screens[home.Identifier] = home
	Screens[login.Identifier] = login
	Screens[register.Identifier] = register
}

// PushScreen adds a new screen to the navigation-stack.
func PushScreen(state *models.State, screen Screen) error {
	navigationStack.Push(screen)
	err := display(state, false)
	return err
}

// PopScreen removes the top screen.
func PopScreen(state *models.State) error {
	navigationStack.Pop()
	err := display(state, false)
	return err
}

// PopAndPushScreen removes the topscreen, adds a new screen, and displays it.
func PopAndPushScreen(state *models.State, screen Screen) error {
	err := PopScreen(state)
	if err != nil {
		return err
	}
	err = PushScreen(state, screen)
	if err != nil {
		return err
	}
	return nil
}

// RemoveScreens clears the navigation stack.
func RemoveScreens(state *models.State) error {
	for i := 0; i < navigationStack.Len(); i++ {
		_, err := navigationStack.Pop()
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveScreensAndPush clears the navigation stack and pushes the screen specified.
func RemoveScreensAndPush(state *models.State, screen Screen) error {
	for i := 0; i < navigationStack.Len(); i++ {
		_, err := navigationStack.Pop()
		if err != nil {
			return err
		}
	}

	navigationStack.Push(screen)
	err := display(state, true)
	return err
}

func display(state *models.State, clear bool) error {
	screen, err := navigationStack.Peek()
	if err != nil {
		return err
	}

	if clear {
		Clear()
	}

	return screen.Show(state)
}
