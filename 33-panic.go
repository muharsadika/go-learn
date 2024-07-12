// dalam golang tidak ada try and catch
// untuk mengatasi error
// menggunakan defer, panic dan recover

package main

import "fmt"

func endApp() {
	fmt.Println("End App")

	message := recover()
	fmt.Println(message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("App Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Hello World")
}
