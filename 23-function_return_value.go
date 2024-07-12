package main

import "fmt"

func sayHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	fmt.Println(sayHello("Muhar Sadika"))

	name := sayHello("Muhar Sadika")
	fmt.Println(name)
}
