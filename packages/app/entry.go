package main

import (
	"fmt"

	"github.com/chriswilliams1977/go-by-example/packages/greet"
	"github.com/chriswilliams1977/go-by-example/packages/greet/de"
)

var (
	a = b
	b = c
	c = 2
	d = x()
)

var integers [10]int

func init() {
	fmt.Println("this is called first")
	//used for global variables
	//for loop not valid for package scope
	for i := 0; i < 10; i++ {
		integers[i] = i
	}
}

func init() {
	fmt.Println("this is called second")
}

func main() {
	fmt.Println("this is called third")
	/*
		var (
			d = e
			e = f
			f = 2
		)
	*/
	//Must export vars in other packages using caps
	fmt.Println(greet.Morning)
	//you can nest packages
	fmt.Println(de.Morning)
	//can be accesses without export (capitals) as its in same package (package scope)
	//must be outside a function
	//Cannot be redeclared in same package
	fmt.Println(version)
	//This compiles as go can work out correct intialization when in package scope
	//c declared first, then b, then a
	fmt.Printf("a is %v \nb is %v \nc is %v \nd is %v\n", a, b, c, d)
	//wont compile as vars are inside function and intialization cycle will not work
	//fmt.Printf("d is %v \ne is %v \nf is %v \n", d, e, f)
	fmt.Println(integers)
}
