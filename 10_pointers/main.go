package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n%T\n",a, b)

	// Use * to read val from address
	fmt.Println(b,"equals",*b)

	// Change val with point
	*b = 10
	fmt.Println(b,"equals",*b)
	fmt.Println(&a,"equals",a)
}
