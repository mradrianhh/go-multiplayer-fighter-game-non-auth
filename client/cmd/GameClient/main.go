package main

import (
	"fmt"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/client/network"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/models"
)

var response int

func main() {
	fmt.Println("Starting client...")

	for {
		fmt.Println("1 - Message Server | 2 - Exit")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				msg, err := network.MessageServer(models.NewMessage(models.AUTHENTICATION, "adrianhh\nfjellhomse123", models.NEUTRAL))
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println(msg.Message)
			case 2:
				os.Exit(0)
			default:
				fmt.Println("Can't understand")
			}
		}
	}
}
