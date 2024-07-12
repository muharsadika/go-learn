package main

import "fmt"

func main() {
	type NoKTP string

	var ktpSadika NoKTP = "11111111"

	var contoh string = "22222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpSadika)
	fmt.Println(contohKtp)
}
