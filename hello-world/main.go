package main

import "fmt"

func main() {
	// Define a variable
	message := "Hello World"
	// Get the reference to the memory location
	// of the message variable (Pointer)
	message_pointer := &message
	// Print the message content via Pointer
	fmt.Println(*message_pointer)
}
