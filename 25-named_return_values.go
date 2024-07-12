package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "Muhar"
	lastName = "Sadika"

	return firstName, lastName
}

func main() {
	firstName, lastName := getCompleteName()
	fmt.Println(firstName, lastName)
}
