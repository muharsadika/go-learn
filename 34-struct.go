package main

import "fmt"

type TCustomer struct {
	Name    string
	Address string
	Age     int
}

func (customer TCustomer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var sadika TCustomer
	sadika.Name = "Muhar Sadika"
	sadika.Address = "Indonesia"
	sadika.Age = 20
	fmt.Println(sadika)
	fmt.Println(sadika.Name)
	fmt.Println(sadika.Address)
	fmt.Println(sadika.Age)

	eko := TCustomer{
		Name:    "Eko",
		Address: "Indonesia",
		Age:     20,
	}
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	budi := TCustomer{"Budi", "Indonesia", 20}
	fmt.Println(budi)
	fmt.Println(budi.Name)
	fmt.Println(budi.Address)
	fmt.Println(budi.Age)

	sadika.sayHello("agus")
	eko.sayHello("agus")
	budi.sayHello("agus")
}
