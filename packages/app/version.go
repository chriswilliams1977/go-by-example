package main

import "fmt"

var version = "1.0.1"

func init() {
	fmt.Println("init func in a package scope file")
}

func x() int {
	return c + 1
}
