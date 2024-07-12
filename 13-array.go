package main

import "fmt"

func main() {
	var names [2]string

	names[0] = "Muhar"
	names[1] = "Sadika"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names)

	var values = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values)

	var newValues = [...]int{
		1,
		2,
		3,
	}

	fmt.Println(newValues)
	fmt.Println(len(newValues))
}
