package main

import "fmt"

type Address struct {
	City, Province, Counry string
}

func main() {
	address1 := Address{"Depok", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Semarang"

	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // operator asterisk

	fmt.Println(address1)
	fmt.Println(address2)

}