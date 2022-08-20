package main

import (
	"fmt"
	"sort"
)

// most used in golang
func main() {
	fmt.Println("Wel Slices")

	var colors = []string{"RED", "BLACK", "BLUE"}
	fmt.Printf("Type %T\n", colors)

	colors = append(colors, "GREEN", "CYAN")
	fmt.Println(colors)

	colors = append(colors[1:]) // range to slice the array source:dest dest is not inclusive
	fmt.Println(colors)

	//array using valours operator
	empid := make([]int, 2)

	empid[0] = 1
	empid[1] = 2
	//empid[2] = 3 //not allowed since initialized size is 2

	empid = append(empid, 4, 3) // go allows to reinitialize the memory using append here and covert as slice now

	fmt.Println(empid)

	sort.Ints(empid) // sort like method is ava in slice not in  array
	fmt.Println(sort.IntsAreSorted(empid))

	// remove a value from slices based on index

	var index int = 2
	colors = append(colors[:index], colors[index+1:]...)
	fmt.Println(colors)

}
