package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
