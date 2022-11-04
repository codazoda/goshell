package main

import (
	"fmt"
)

func main() {
	var name string
	// Loop forever
	for {
		// Ask the user for their name
		fmt.Print("Who goes there? ")
		// Wait for the users input
		fmt.Scanln(&name)
		// If the user is exit, quit
		if name != "exit" {
			// Say hello
			fmt.Println("Hello " + name)
		} else {
			// Bail out of the forever loop
			return
		}
	}
}
