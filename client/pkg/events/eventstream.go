package events

// GlobalEventStream is a channel where one can push events into one end, and subscribers can receive them on the other.
var GlobalEventStream chan string = make(chan string)

// Notify notifies all listeners on the event stream.
func Notify(eventStream chan<- string, event string) {
	eventStream <- event
}
