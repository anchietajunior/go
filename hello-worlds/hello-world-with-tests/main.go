package main

import (
	"fmt"
)

func printMessage(message string) string {
	message_pointer := &message
	return *message_pointer
}

func main()  {
	fmt.Println(printMessage("Hello World"))
}

