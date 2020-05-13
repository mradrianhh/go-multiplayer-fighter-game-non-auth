package navigation

import "github.com/mradrianhh/go-multiplayer-fighter-game/internal/pkg/models"

// Screen ...
type Screen interface {
	Show(state *models.State) error
}

// Empty ...
type Empty struct {
}

// Show ...
func (empty Empty) Show(state *models.State) error {
	return nil
}
