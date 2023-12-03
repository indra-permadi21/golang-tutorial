package main

import "fmt"

func main() {
	const (
		firstName string = "Indra"
		lastName         = "Permadi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// constant can't change the value
	//firstName = "Dodit"
	//lastName = "Mulyadi"
}
