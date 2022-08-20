package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	wel := "Wel to user input"
	fmt.Println(wel)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the vaule:")

	//commo ok || err err syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Value:", input)
	fmt.Printf("Type %T", input)
}
