// Package network concerns itself with the communications between the game-client and the game-server.
//
// The network-communications system is designed in such a way, that a connection between the client and the server will never be maintained over longer amounts of time.
// 	Instead, the client makes sure the user is authenticated and logged in. As the client initializes the game-state to "HOME",
//	the user will never be able to access the game without logging in through the "login"-screen. Once the user logs in, he's directly brought to the main-menu.
//	As the state is never recorded, the second he passes the "login"-screen, the state has changed to "MAINMENU", and all notions of the user ever being authenticated is thrown out the window.
//	If the user ever logs out, the game will transfer to the "HOME"-state, and the user will again have to log in to be able to reach the main-menu.
//
// Proper security is not dealt with. The engine's communication system is not meant, nor designed, for transferring private data.
//	The whole idea of a "username" and "password", is for the players to be able to identify themselves and retrieve their saves.
//	As such, when selecting a username and password in the register screen, do not pass in a private password that you're using in any other regard.
//	The password is sent to the server in plain text, is saved in plain text, and thus, it cannot be stressed enough: none of the information is encrypted, safe or private.
package network
