package main

import "fmt"

func main() {
	name := "Sadika"

	switch name {
	case "Muhar":
		fmt.Println("Hello Muhar")
	case "Sadika":
		fmt.Println("Hello Sadika")
	default:
		fmt.Println("Hi, my name is", name)
	}
}
