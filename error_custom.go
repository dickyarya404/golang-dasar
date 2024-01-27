package main

import (
	"fmt"
)

type validationEror struct { // custom error karena mengikuti kontrak error
	Message string
}

func (e *validationEror) Error() string {
	return e.Message
}

type notFoundError struct { // custom error karena mengikuti kontrak error
	Message string
}

func (e *notFoundError) Error() string {
	return e.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationEror{"validation error"}
	}

	if id != "Joe" {
		return &notFoundError{"Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("Joe  ", nil)
	if err != nil {
		// terjadi error
		// if validationErr, ok := err.(*validationEror); ok { // untuk mengecek apakah bisa di koversi atau tidak
		// 	fmt.Println("validation error:", validationErr.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("Unknown error:", err.Error())
		// }

		// menggunakan swict case
		switch finalError := err.(type) {
			case *validationEror:
				fmt.Println("validation error:", finalError.Error())
			case *notFoundError:
				fmt.Println("not found error:", finalError.Error())
			default:
				fmt.Println("Unknown error:", err.Error())
		}	


	} else {
		// Sukses
		fmt.Println("sucess")

	}

}