package main

import "fmt"

// array is less used in golang

func main() {
	fmt.Println("Wel Array")

	var arr [5]string

	arr[0] = "num1"
	arr[1] = "num2"
	arr[3] = "num3"

	fmt.Println("Array value:", arr)
	fmt.Println("Array value:", len(arr)) // i have inserted the 3, output will be 5 since i have reversed the memory

	var arr2 = [3]string{"num1", "num1", "num1"}
	fmt.Println("Array value:", arr2)
	fmt.Println("Array value:", len(arr2))
}
