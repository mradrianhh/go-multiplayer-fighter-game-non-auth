package navigation

import "github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/models"

// Navigator holds a stack, and a map of screens to use in the stack. Through it's method, one may manipulate the stack, hence altering the appearance of the screen.
type Navigator struct {
	stack   Stack
	screens map[string]Screen
}

// NewNavigator returns a new Navigator-instance with default settings.
func NewNavigator() Navigator {
	var navigator Navigator
	navigator.screens = make(map[string]Screen)
	return navigator
}

// PushScreen adds a new screen to the navigation-stack.
func (navigator *Navigator) PushScreen(state *models.State, identifier string) error {
	navigator.stack.Push(navigator.screens[identifier])
	err := navigator.display(state, false)
	return err
}

// PopScreen removes the top screen.
func (navigator *Navigator) PopScreen(state *models.State) error {
	navigator.stack.Pop()
	err := navigator.display(state, false)
	return err
}

// PopAndPushScreen removes the topscreen, adds a new screen, and displays it.
func (navigator *Navigator) PopAndPushScreen(state *models.State, identifier string) error {
	err := navigator.PopScreen(state)
	if err != nil {
		return err
	}
	err = navigator.PushScreen(state, identifier)
	if err != nil {
		return err
	}
	return nil
}

// RemoveScreens clears the navigation stack.
func (navigator *Navigator) RemoveScreens(state *models.State) error {
	for i := 0; i < navigator.stack.Len(); i++ {
		_, err := navigator.stack.Pop()
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveScreensAndPush clears the navigation stack and pushes the screen specified.
func (navigator *Navigator) RemoveScreensAndPush(state *models.State, identifier string) error {
	for i := 0; i < navigator.stack.Len(); i++ {
		_, err := navigator.stack.Pop()
		if err != nil {
			return err
		}
	}

	navigator.stack.Push(navigator.screens[identifier])
	err := navigator.display(state, true)
	return err
}

// AddScreen assigns a new screen to the map with it's identifier as the key.
func (navigator *Navigator) AddScreen(identifier string, screen Screen) {
	navigator.screens[identifier] = screen
}

func (navigator *Navigator) display(state *models.State, clear bool) error {
	screen, err := navigator.stack.Peek()
	if err != nil {
		return err
	}

	if clear {
		Clear()
	}

	return screen.Show(state)
}
