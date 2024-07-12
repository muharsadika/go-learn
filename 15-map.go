package main

import "fmt"

func main() {
	// var person = map[string]string{}
	// person["name"] = "Sadika"
	// person["title"] = "Programmer"

	var person = map[string]string{
		"name":  "Sadika",
		"title": "Programmer",
	}

	fmt.Println(person["name"])
	fmt.Println(person["title"])
	fmt.Println(person)

	book := make(map[string]string)

	book["title"] = "Belajar Golang"
	book["author"] = "Muhar Sadika"
	book["ups"] = "Salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
