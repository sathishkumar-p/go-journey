package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	wel := "Wel to user input"
	fmt.Println(wel)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the vaule:")

	//commo ok || err err syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Value:", input)
	fmt.Printf("Type %T \n", input)

	numVal, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// input read both value and \n

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 1:", numVal+1)
	}
}
