package main

import (
	"fmt"
)

func main() {

	name := "Dicky"

	switch name {
	case "Dicky":
		fmt.Println("Hello Dicky")
	case "Jho":
		fmt.Println("Hello Jho")
	default:
		fmt.Println("Hallo bro")

	}

	switch length := len(name); length > 5 {
		case true:
			fmt.Println("Nama terlalu panjang")
		case false:
			fmt.Println("Nama sudah benar")	
	}

	// switch pengganti if else
	name = "DickyAryaAjiPrayoga"

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama agak panjang")
	default:
		fmt.Println("Nama sudah benar")	
	}

}   