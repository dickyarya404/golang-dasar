package main

import "fmt"

// type assertion merupaakn kemampuan merubah tipe data yg di inginkan dengan tipe data kosong

func random() bool {
	return true
}


func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	} 

}