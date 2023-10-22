package main

import "fmt"

func main() {
	// -------------------------------//
	// Run Quiz Method
	var car map[string]string
	car = map[string]string{}
	car["name"] = "BWM"
	car["color"] = "Black"

	printString(car["name"], car["color"])
	// output => Mobil BMW berwarna Black

	// -------------------------------//
	// Run Quiz Pointer
	student := Student{
		Name:  "Nobee",
		Class: "Intermediate",
	}
	student.CallMyName()
	//	output => Hello, My Name is Nobee I'm already on class  Intermediate
}

// Quiz Method
func printString(name string, color string) {
	fmt.Println("Mobil", name, "berwarna", color)
}

// Quiz Pointer
type Student struct {
	Name  string
	Class string
}

func (s *Student) SetMyName(name string) {
	s.Name = name
}

func (s *Student) CallMyName() {
	fmt.Println("Hello, my name is", s.Name, "i'm already on class", s.Class)
}
