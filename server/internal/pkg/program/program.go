package program

import (
	"fmt"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/server/internal/pkg/network"
)

var response int

// Run starts the program.
func Run() {
	ShowMainMenu()
}

// ShowMainMenu presents the user with the main-menu.
func ShowMainMenu() {
	for {
		fmt.Println("\nMain Menu\n\n1 - View all messages\n2 - View log\n3 - View all errors\n4 - View players\n5 - View connections\n6 - View authenticated connections\n0 - Exit")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				ShowMessages()
			case 2:
				ShowLog()
			case 3:
				ShowErrorLog()
			case 4:
				ShowPlayers()
			case 5:
				ShowConnections()
			case 6:
				ShowAuthenticatedConnections()
			case 0:
				Clear()
				os.Exit(0)
			default:
				fmt.Println("Sorry, I can't understand...")
			}
		}
	}
}

// ShowMessages presents the user with all the recorded messages in the service.
func ShowMessages() {
	Clear()
	fmt.Print("Messages\n\n")
	if len(network.Messages) > 0 {
		for index, value := range network.Messages {
			value.Present(index + 1)
		}
	} else {
		fmt.Println("No messages to show.")
	}
}

// ShowLog presents the user with the log of the service.
func ShowLog() {
	Clear()
	fmt.Print("Log\n\n")
	if len(network.Log) > 0 {
		for index, value := range network.Log {
			fmt.Printf("%v. %s", index, value)
		}
	} else {
		fmt.Println("No log to show.")
	}
}

// ShowErrorLog presents the user with the error-log of the service.
func ShowErrorLog() {
	Clear()
	fmt.Print("Error Log\n\n")

	if len(network.ErrorLog) > 0 {
		for index, value := range network.ErrorLog {
			fmt.Printf("%v. %s\n", index+1, value.Error())
		}
	} else {
		fmt.Println("No errors to show.")
	}
}

// ShowPlayers presents the user with a list of all the players connected to the service.
func ShowPlayers() {
	Clear()
	fmt.Print("Players\n\n")

	if len(network.Players) > 0 {
		index := 1
		for _, value := range network.Players {
			value.Present(index)
			index++
		}
	} else {
		fmt.Println("No players connected.")
	}
}

// ShowConnections presents the user with a list of all the connections to the server.
func ShowConnections() {
	Clear()
	fmt.Print("Connections\n\n")

	if len(network.Conns) > 0 {
		index := 1
		for _, conn := range network.Conns {
			fmt.Printf("%v. %s.\n", index, conn.LocalAddr().String())
			index++
		}
	} else {
		fmt.Println("No connections to show.")
	}
}

// ShowAuthenticatedConnections presents the user with a list of all the authenticated connections to the server.
func ShowAuthenticatedConnections() {
	Clear()
	fmt.Print("Authenticated connections\n\n")

	if len(network.AuthenticatedConns) > 0 {
		index := 1
		for key, value := range network.AuthenticatedConns {
			fmt.Printf("%v. %s - %s.\n", index, key, value.LocalAddr().String())
			index++
		}
	} else {
		fmt.Println("No connections to show.")
	}
}
