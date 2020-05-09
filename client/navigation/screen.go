package navigation

import "github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"

// Screen ...
type Screen interface {
	Show(state *vars.State) error
}

// Empty ...
type Empty struct {
}

// Show ...
func (empty Empty) Show(state *vars.State) error {
	return nil
}
