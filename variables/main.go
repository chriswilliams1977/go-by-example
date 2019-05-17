package main

import (
	"fmt"
)

const(
    a = iota // a == 0
    b = iota // b == 1
	c = iota // c == 2
	_		 // skipped value
	d        // d == 3 (implicitely d = iota)
)

var intZeroValue int
type  myStringObject string

func main() {
	var myString myStringObject = "new string type"
	fmt.Printf("The zero value is %v \n" , intZeroValue)
	fmt.Printf("The value of myString is %v and the type is %T \n", myString, myString)
	fmt.Printf("Iota incremented value is %v \n", d)
}