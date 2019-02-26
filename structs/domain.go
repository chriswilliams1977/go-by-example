package main

import (
	"fmt"
)

//STRUCTS

//Person : person struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Mother    *Mother
}

//Mother : composition example. Inherits properties and methods from parent
type Mother struct {
	Person
}

//Dog : New struct
type Dog struct {
}

//INTERFACE

//lifeform : Defines an interface if used all subscribers must implement talk method
type lifeform interface {
	talk()
}

//INTERFACE METHODS

func (p *Person) talk() {
	fmt.Printf("My name is %v and I am %v years old. \r\n", p.FirstName+" "+p.LastName, p.Age)
}

func (d *Dog) talk() {
	fmt.Printf("Woof woof \r\n")
}

func doSomething(lf lifeform) {
	lf.talk()
}
