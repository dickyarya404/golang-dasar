package main

import "fmt"

func main() {
	type NoKtp string // tipe dara nama lainnya

	var ktpDik NoKtp = "21490u9413431412"

	var contoh string = "222222222222222"
	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(ktpDik)
	fmt.Println(contohKtp)
}