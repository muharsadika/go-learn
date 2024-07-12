package main

import "fmt"

type TFilter func(string) string

func sayHelloWithFilter(name string, filter TFilter) {
	nameFiltered := filter(name)

	fmt.Println("Hello", nameFiltered)
}

func filter(name string) string {
	if name == "Anjing" {
		return "****"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing", filter)
}
