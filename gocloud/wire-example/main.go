package main

import (
	"fmt"
	"os"
)

func main() {

	//Set up without wire
	//Have to set up all initialization first then execute
	//We are using the dependency injection design principle.
	//In practice, that means we pass in whatever each component needs.
	//Downside with DI is all the initialization
	//Use wire to simplify
	/*
		message := NewMessage()
		greeter := NewGreeter(message)
		event := NewEvent(greeter)
	*/

	e, err := InitializeEvent("Hello there!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()

}
