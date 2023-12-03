package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	var (
		firstName  = "indra"
		middleName = "No Middle Name"
		lastName   = "Permadi"
	)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}
