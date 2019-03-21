package main

//Message : message for greeter
type Message string

//NewMessage : message initializer
//returns a message with string set to phrase
func NewMessage(phrase string) Message {
	return Message(phrase)
}
