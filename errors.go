package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) { // tipe interface error 
	if pembagi == 0 {
		return 0, errors.New("Pemabagi dengan nol ") // membuat error gunakan Error.New
	} else {
		return nilai / pembagi, nil
	}
	
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil: ", hasil)
	} else {
		fmt.Println("Error: ", err.Error())
	}

}