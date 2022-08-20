package main

import "fmt"

func main() {
	fmt.Println("Wel Pointers")

	// var myptr *int
	// fmt.Println("Value of pointer is ", myptr)
	//reference means &

	num := 23

	var pt = &num
	fmt.Println("Address vaule", pt)
	fmt.Println("Actual vaule", *pt)

	*pt = *pt * 2
	fmt.Println("New value", num)
}
