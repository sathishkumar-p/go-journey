package main

import "fmt"

func main() {
	fmt.Println("Loops")

	values := []string{"1", "2", "3", "4", "5"}

	fmt.Println(values)

	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
	}

	for key := range values {
		fmt.Println(values[key])
	}
	for k, v := range values {
		fmt.Printf("Key is %v and value is %v\n", k, v)
	}

	rangeValue := 1

	for rangeValue < 10 {
		if rangeValue == 2 {
			rangeValue++
			continue
			//break
		}
		if rangeValue == 5 {
			goto some
		}
		fmt.Println("value is", rangeValue)
		rangeValue++
	}
	fmt.Println("Goto skipped")
some:
	fmt.Println("Goto")
}
