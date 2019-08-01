package main

import (
	"fmt"
)

func main() {
	ids := []int{33,67,54,32,33,2,11,67}
	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com", "Mike":"mike@gmail.com"}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i ,id)
	}
	for _, id := range ids {
		fmt.Printf("ID: %d\n",id)
	}

	sum := 0
	for _, id := range ids {
		sum+= id
	}
	fmt.Println("Sum", sum)

	// Range with map
	for k,v := range emails {
		fmt.Printf("%s: %s\n", k,v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
