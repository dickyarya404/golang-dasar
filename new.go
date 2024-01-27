package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func main() {
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat1.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}