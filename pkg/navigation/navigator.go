package navigation

import "github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/models"

var navigationStack Stack

// Screens is a key-value collection of the screens registered.
// Add values to the map by assigning a screen with the screen's identifier as the key.
var Screens map[string]Screen

func init() {
	Screens = make(map[string]Screen)
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
