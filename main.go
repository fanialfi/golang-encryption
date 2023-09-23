package main

import (
	"fmt"
	"golang-encryption/lib"
)

func main() {
	text := "ini rahasia"
	rahasiaBumi := lib.HashingSederhana(text)
	fmt.Println(rahasiaBumi)

	rahasia := lib.HashingSederhanaV2(text)
	fmt.Println(rahasia)
	fmt.Println()

	hashedPassword1 := lib.Salting(text)
	fmt.Println(hashedPassword1)
	fmt.Println()

	hashedPassword2 := lib.Salting(text)
	fmt.Println(hashedPassword2)
	fmt.Println()

	hashedPassword3 := lib.Salting(text)
	fmt.Println(hashedPassword3)
}
