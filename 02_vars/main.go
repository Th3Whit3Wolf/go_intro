package main

import (
	"fmt"
)

func main() {
	// MAIN TYPES
	//	string
	//	bool
	//	int
	//	int		int8 	int16	int32	int64
	//	uint	uint8	uint16	unint32	unint64	unintptr
	//	byte -	alias for uint8
	//	rune -	alias for int32
	//	float32	float64
	//	complex64	complex128

	//Shorthand
	//name := "David"
	//email := "david@mail.com"

	name, email := "David", "david@mail.com"

	var age uint8 = 27
	var weight float32 = 181.2
	const isCool = true

	fmt.Println(name, "is", age, "years old and weighs", weight, "lbs")
	fmt.Print("and you can reach him at ", email)
	fmt.Printf(" %T\n", isCool)
}
