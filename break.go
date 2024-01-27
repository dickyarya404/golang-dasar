package main

import "fmt"

// break
func main() {
	for i := 0; i < 10; i++ {
		if  i == 5 { 
			break // menghentikan perulangan dengan break
		}
		fmt.Println("Perulangan ke", i)
	}

}