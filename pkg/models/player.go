package models

// Players is the global collection of all players in-game.
var Players []Player

// Player represents the user in-game.
type Player struct {
	Username string
	Password string
	Online   bool
	Exp      int
}

// NewPlayer works as a constructor, creating a new Player-object with the specified username attribute.
func NewPlayer(username string, online bool, password string) Player {
	return Player{Username: username, Exp: 0, Online: online, Password: password}
}

// UpdateExp alters the player's exp.
func (player *Player) UpdateExp(exp int) {
	player.Exp += exp
}

// SetOnline alters the player's online status.
func (player *Player) SetOnline(online bool) {
	player.Online = online
}
