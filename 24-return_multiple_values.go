package main

import "fmt"

func getFullName() (string, string) {
	return "Muhar", "Sadika"
}

func getFullName2() (string, string) {

	return "Muhar", "Sadika"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	firstName2, _ := getFullName2()
	fmt.Println(firstName2)
}
