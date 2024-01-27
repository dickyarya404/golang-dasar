package main

import "fmt"

func main() {
	name := "Prayoga" // menggunakan tipe boolean true or false

	if name == "Dicky" {
		fmt.Println("Hello Dicky!")
	} else if name == "Joko"{
		fmt.Println("Hello, Joko")
	} else {
		fmt.Println("Hi, Apa kabar!")
	}

	
	// if statment
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}