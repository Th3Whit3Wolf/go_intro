package main

import "fmt"

func main() {
	// Define Map
	// emails := make(map[string]string)

	// Assign KV
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com", "Mike":"mike@gmail.com"}
	fmt.Println(emails)
	fmt.Println(emails["Sharon"])
	fmt.Println(len(emails))

	// Delete from Map
	delete(emails, "Bob")
	fmt.Println("\nBob has been removed from the map\n")

	fmt.Println(emails)
	fmt.Println(emails["Sharon"])
	fmt.Println(len(emails))
}
