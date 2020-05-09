package events

import "fmt"

// EventListener listens for events.
type EventListener struct {
	Events map[string]func()
}

// GlobalEventListener ...
var GlobalEventListener EventListener

func init() {
	GlobalEventListener = EventListener{}
	GlobalEventListener.Events = make(map[string]func())
}

// Listen loops through all the events of the eventListener and checks the eventStream for any occurences. If there are any, it executes the associated function.
func (EventListener *EventListener) Listen(eventStream <-chan string) {
	for {
		if event, ok := <-eventStream; ok {
			for e, f := range EventListener.Events {
				if event == e {
					fmt.Println(event)
					f()
				}
			}
		}
	}
}
