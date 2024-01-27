package main

import (
	"fmt"
	"golang-dasar/database"
	_ "golang-dasar/internal" // _blank indentifire untuk mengesekusi kode init saja
)

func main() {
	fmt.Println(database.GetDatabase())
}


