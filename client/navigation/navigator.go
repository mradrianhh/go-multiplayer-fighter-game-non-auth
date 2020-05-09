package navigation

import (
	"github.com/mradrianhh/go-multiplayer-fighter-game/client/screens"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
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
func PushScreen(state *vars.State, screen Screen) error {
	navigationStack.Push(screen)
	err := display(state)
	return err
}

// PopScreen removes the top screen. If displayPrev is true, it will also reveal the now-topmost screen.
func PopScreen(state *vars.State, displayPrev bool) error {
	navigationStack.Pop()
	if displayPrev {
		err := display(state)
		return err
	}
	return nil
}

// PopAndPushScreen removes the topscreen, adds a new screen, and displays it.
func PopAndPushScreen(state *vars.State, screen Screen) error {
	err := PopScreen(state, false)
	if err != nil {
		return err
	}
	err = PushScreen(state, screen)
	if err != nil {
		return err
	}
	return nil
}

func display(state *vars.State) error {
	screen, err := navigationStack.Peek()
	if err != nil {
		return err
	}

	Clear()
	err = screen.Show(state)
	if err != nil {
		return err
	}
	return nil
}

// RemoveScreens clears the navigation stack.
func RemoveScreens(state *vars.State) error {
	for i := 0; i < navigationStack.Len(); i++ {
		err := PopScreen(state, false)
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveScreensAndPush clears the navigation stack and pushes the screen specified.
func RemoveScreensAndPush(state *vars.State, screen Screen) error {
	for i := 0; i < navigationStack.Len(); i++ {
		err := PopScreen(state, false)
		if err != nil {
			return err
		}
	}

	err := PushScreen(state, screen)
	return err
}
