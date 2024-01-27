package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { // tambahkan * untuk jadi pointerm 
	man.Name = "Mr. " + man.Name
}

func main() {
	joe := Man{"joe"}
	joe.Married()

	fmt.Println(joe.Name)

}