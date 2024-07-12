package main

import "fmt"

func main() {
	// var nilai32 int32 = 3278
	// var nilai64 int64 = int64(nilai32)
	// var nilai16 int16 = int16(nilai32)

	// fmt.Println(nilai32)
	// fmt.Println(nilai64)

	// // hasil dari nilai16 akan bernilai -32768
	// // karena int16 mempunyai range -32768 sampai 32767
	// // jika melebihi batas, maka angka akan mutar kembali ke angka paling kecil di int16, yaitu -32768
	// fmt.Println(nilai16)
	var name = "Muhar Sadika"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
