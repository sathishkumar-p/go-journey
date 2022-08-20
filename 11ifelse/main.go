package main

import "fmt"

func main() {
	fmt.Println("If else")

	num := 3

	if num > 10 {
		fmt.Println("Number is > 10")
	} else if num < 10 {
		fmt.Println("Number is < 10")
	} else {
		fmt.Println("Number is = 10")
	}

	if 7%3 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if val := 10; val > 10 {
		fmt.Println("Number is > 10")
	} else {
		fmt.Println("Number is = 10")
	}
}
