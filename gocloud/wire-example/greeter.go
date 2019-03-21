package main

import (
	"time"
)

//Greeter : Say  message
//Property of Message
type Greeter struct {
	Message Message
	Grumpy  bool
}

//NewGreeter : Greeter initializer
//pass message and set prop to message value
func NewGreeter(m Message) Greeter {
	// set bool
	var grumpy bool
	//if the invocation time of the initializer is an even number of seconds
	//we will create a grumpy greeter instead of a friendly one.
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

//Greet : Greet method on Greeter
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}
