package navigation

import "errors"

// Stack errors.

// ErrEmptyStack is thrown when the user is attempting to retrieve an item from the stack, but it's empty.
var ErrEmptyStack = errors.New("stack is empty")
