package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) { // asterik
	address.Country = "Indonesia"
}

func main() {
	var address *Address = &Address{} // jadikan pointer
	ChangeCountryToIndonesia(address)

	fmt.Println(address)
}