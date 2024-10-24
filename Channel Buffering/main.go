package main

import "fmt"

func main() {
	// channel string buffering up to 2 values
	messages := make(chan string, 2)
	
	messages <- "value1"
	messages <- "value2"
	// here we are able to send these 2 vals since we created a cuffering for 2 values above
	
	
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}