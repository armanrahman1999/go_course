package main

import (
	"fmt"
	"time"
)

func sendMessage(num int) {
	fmt.Printf("Sending message %d\n", num)

	time.Sleep(time.Second * time.Duration(num)) // Simulate some work

	msg := fmt.Sprintf("✅ Message %d sent!", num)
}

func main() {

}