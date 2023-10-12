package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	result := make(chan string)

	go process1(message, result)
	go process2(message)

	fmt.Println(<-result)

}

func process1(message chan string, result chan string) {
	fmt.Println("Running proccess 1")
	time.Sleep(1 * time.Second)

	fmt.Println("Process 1: Wait until got message form proccess 2")
	m := <-message
	fmt.Println("Process 1: Got message from ", m)

	result <- "Done from process 1"
}

func process2(message chan string, result chan string, done chan bool) {
	fmt.Println("Running process 2")
	time.Sleep(1 * time.Second)
	fmt.Println("Process 2: Send message to channel message")
	message <- "From process 2"

	fmt.Println("Process 2: wait from result")
	res := <-result

	fmt.Println(res)
	done <- true

}
