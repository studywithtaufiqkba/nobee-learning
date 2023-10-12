//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	// membuat sebuah channel string
//	process := make(chan string)
//
//	// masukin ke goroutine
//	go process1(process)
//	go process2(process)
//	go process3(process)
//
//	// print value
//	fmt.Println("value :", <-process)
//	fmt.Println("value :", <-process)
//	fmt.Println("value :", <-process)
//}
//
//func process1(process chan string) {
//	msg := "Process 1"
//	fmt.Println("process running :", msg)
//
//	// proses memasukkan sebuah value ke dalam channel
//	process <- msg
//}
//func process2(process chan string) {
//	msg := "Process 2"
//	fmt.Println("process running :", msg)
//
//	// proses memasukkan sebuah value ke dalam channel
//	process <- msg
//}
//func process3(process chan string) {
//	msg := "Process 3"
//	fmt.Println("process running :", msg)
//
//	// proses memasukkan sebuah value ke dalam channel
//	process <- msg
//}
