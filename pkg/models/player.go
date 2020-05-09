package models

// Players is the global collection of all players in-game.
var Players []Player

// Player represents the user in-game.
type Player struct {
	Username string
	Exp      int
}

// NewPlayer works as a constructor, creating a new Player-object with the specified username attribute.
func NewPlayer(username string) Player {
	return Player{Username: username, Exp: 0}
}

// UpdateExp alters the player's exp.
func (player *Player) UpdateExp(exp int) {
	player.Exp += exp
}
