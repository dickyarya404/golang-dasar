package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your Are Blocked", name)
	} else {
		fmt.Println("Welcome to", name) 
	}
}

func main() {
	blacklist := func (name string) bool {
		return name == "toxic"
	}

	registerUser("Dicky", blacklist) // func blacklist

	registerUser("toxic", func(name string) bool { // jika ada parameter yang butuh lansung menggunaka 
		// anonymous function
		return name == "toxic"
	}) 


}