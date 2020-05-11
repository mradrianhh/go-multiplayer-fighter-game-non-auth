package models

import "fmt"

// Players is the global collection of all players in-game.
var Players []Player

// Player represents the user in-game.
type Player struct {
	Username string
	Password string
	Exp      int
}

// NilPlayer is the empty player value.
var NilPlayer = Player{}

// NewPlayer works as a constructor, creating a new Player-object with the specified username attribute.
func NewPlayer(username string, password string) Player {
	return Player{Username: username, Exp: 0, Password: password}
}

// UpdateExp alters the player's exp.
func (player *Player) UpdateExp(exp int) {
	player.Exp += exp
}

// Present prints the player to the screen.
func (player Player) Present(index int) {
	fmt.Printf("Player %v\n", index)
	fmt.Printf("\tUsername: %s\n", player.Username)
	fmt.Printf("\tPassword: %s\n", player.Password)
	fmt.Printf("\tExp: %v\n\n", player.Exp)
}
