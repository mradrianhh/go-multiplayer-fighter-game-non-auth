package main

import (
	"github.com/mradrianhh/go-multiplayer-fighter-game/client/internal/pkg/program"
	"github.com/mradrianhh/go-multiplayer-fighter-game/client/internal/pkg/screens"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/navigation"
)

func main() {
	initialize()

	program.Run()
}

func initialize() {
	mainmenu := screens.GetMainMenu()
	navigation.Screens[mainmenu.Identifier] = mainmenu

	login := screens.GetLogin()
	navigation.Screens[login.Identifier] = login

	register := screens.GetRegister()
	navigation.Screens[register.Identifier] = register

	home := screens.GetHome()
	navigation.Screens[home.Identifier] = home
}
