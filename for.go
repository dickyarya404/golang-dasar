package main

import "fmt"

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }
	// fmt.Println("selesai")

	for counter := 1; counter <= 10; counter++ { // lebih sederhana
		fmt.Println("Perulangan ke", counter)
	}
	fmt.Println("selesai")

	names := []string{"Dicky", "Arya", "Aji"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// kode program for range
	for index, name := range names {
		fmt.Println("index: ", index, "=", name)
	}

	// tanpa index for
	for _, name := range names {
		fmt.Println(name)
	}
}