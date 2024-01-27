package main

import "fmt"

// nil merupakan data kosong
// nil hanya di gunakan pd tipe data sperti interface, func, slice pointer dan chanel
func newMap(name string) map[string]string {
	if name == "" {
		return nil // data kosong
	} else {
		return map[string]string {
			"name": name,
		}
	}

}

func main() {
	data := newMap("Data ready")
	if data == nil {
		fmt.Println("No data")
	} else {
		fmt.Println(data["name"])
	}

	

}