package main

import "fmt"

func main() {
	name := "Sadika"

	if name == "Muhar" {
		fmt.Println("Hello Muhar")
	} else if name == "Sadika" {
		fmt.Println("Hello Sadika")
	} else {
		fmt.Println("Hi, my name is", name)
	}
}
