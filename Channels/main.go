package main

import "fmt"

func main() {
	// Create a new channel with make(chan val-type)
	messages := make(chan string)

	// channel <-: send a val into a channel
	// sending "ping" to messages channel from a new go routine
	go func() { messages <- "ping" }()

	// <- channel: receive a value from the channel
	msg := <-messages
	fmt.Println(msg)
}
