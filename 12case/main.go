package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case")

	rand.Seed(time.Now().UnixNano())

	diceNu := rand.Intn(6) + 1
	fmt.Println("Value of rand is ", diceNu)

	switch diceNu {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("value is 2")
		fallthrough // exe nxt case
	case 3:
		fmt.Println("value is 3")
	default:
		fmt.Println("Default", diceNu)
	}
}
