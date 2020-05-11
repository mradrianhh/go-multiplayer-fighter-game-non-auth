// Package navigation maintains and handles the screens in the game.
//
// Declarations:
// 	- Throughout this engine, the word "screen" will be  used for the collection of text which is presented
// 		at the same time and concerns the same aspect of the game.
// 		I.e: the main-menu screen may write "1 - Start Game | 2 - Exit Game" to the console, and then wait for the user to respond.
//	 	This - including the user's response - is considered a screen.
//
// It presents a stack which can be manipulated through the following functions:
// 	- Pop: removes the topmost screen.
// 	- Push: adds a new screen on top.
// 	- Peek: returns the topmost screen without removing it.
// 	- IsEmpty: checks if the stack is empty and returns true or false based on result.
// 	- Len: returns the amount of elements in the stack.
//
// The package holds a public variable of type stack called "navigationStack" which carries the screens.
// The "navigationStack" holds public wrapper functions around the functions of it's type, Stack, for manipulating it:
// 	- PushScreen: pushes the screen onto the "navigationStack", then displays it.
//		The screen is written without clearing previous screens.
//	- PopScreen: pops the topmost screen of the stack, then displays the screen that is now on top.
//		The screen is written without clearing previous screens.
// 	- PopAndPushScreen: pops the topmost screen of the stack, pushes a new one on, and presents this one.
//		The screen is written without clearing previous screens.
// 	- RemoveScreens: clears the stack and the previous screens. As there are no screens to show now,
//		the user will be presented with a blank console.
// 	- RemoveScreensAndPush: clears the stack and the previous screens, pushes a new one on, and presents this one.
//		As all the previous screens have been cleared, this will be the only thing written to the console.
//
// The navigation package also presents a "Screen"-interface containing a "Show"-method.
// 	Thus, everything that implements a show-method is a screen.
package navigation
