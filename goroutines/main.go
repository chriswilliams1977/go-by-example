package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello World!")
}

//Go acheives concurrency through goroutines
//program runs in a goroutine called main
//and manages data returned from these via channels
func main() {
	fmt.Println("execution started")

	//adds a goroutine
	//executes in background and does not block program execution flow
	go printHello()

	//this will delay execution of last command thus print out Hello Woirld first
	//In reality we do not know hom long a routine will take to execute
	//we need a way to get data back from rountine when executed = Channels
	time.Sleep(10 * time.Millisecond)
	fmt.Println("execution ended")
}
