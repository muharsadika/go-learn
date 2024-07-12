package main

import "fmt"

func factorialLoop(n int) int {
	result := 1

	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursion(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorialRecursion(n-1)
	}
}

func main() {
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursion(5))
}
