package main

import "fmt"

func main() {
	name := "Dicky Arya" // langsung inisialisasi datanya
	// mengunakan := untuk membuta variable
	fmt.Println(name)

	name = "Dicky Aji"
	fmt.Println(name)

	var (
		firstName = "Dicky"
		lastName = "Prayoga"
		
	)

	fmt.Println(firstName, lastName)
}