package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	days := make(map[string]string)

	days["1"] = "Monday"
	days["2"] = "Tuesday"
	days["3"] = "Wednesday"

	fmt.Println(days)
	fmt.Println("Days :", days["1"])

	delete(days, "1")
	fmt.Println(days)
	for key, value := range days {
		fmt.Println(key, value)
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
