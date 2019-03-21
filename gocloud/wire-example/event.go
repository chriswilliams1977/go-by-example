package main

import (
	"errors"
	"fmt"
)

//Event : starts to greet guests
type Event struct {
	Greeter Greeter
}

//NewEvent : event initializer
//Return an error with the event
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

//Start : Starts an event
//tells the greeter to issue a greeting and then prints that message to the screen.
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
