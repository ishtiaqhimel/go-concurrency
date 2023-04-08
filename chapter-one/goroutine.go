package main

import (
	"fmt"
	"time"
)

func main() {
	go helloworld()

	// Question Part
	go func(message string) {
		time.Sleep(2 * time.Second)
		fmt.Println(message)
	}("Ishtiaq")
	// will it be executed?

	time.Sleep(1 * time.Second)
	goodbye()
}

func helloworld() {
	fmt.Println("Hello World")
}

func goodbye() {
	fmt.Println("Good Bye!")
}

/*
The main goroutine starts. Then it invokes the helloworld() goroutine and it starts.
After the helloworld goroutine finishes its operation, the main goroutine waits for 1s.
Then invokes goodbye function.

If we omit Sleep function then main will exit before helloworld to finish its execution.
Thats how we know it was running in different thead!!!
*/
