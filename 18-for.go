package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	fmt.Println("Selesai")

	names := []string{"Muhar", "Sadika", "Eko"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println("Index", index, "=", value)
	}
}
