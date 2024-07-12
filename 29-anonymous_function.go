package main

import "fmt"

type TBlacklist func(string) bool

func registerUser(name string, blacklist TBlacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("sadika", blacklist)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
