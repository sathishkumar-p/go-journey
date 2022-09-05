package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")

	content := "Sample content"

	file, err := os.Create("./test.txt")

	checkErr(err)

	leng, err := io.WriteString(file, content)
	checkErr(err)
	fmt.Println("Length is:", leng)
	defer file.Close() //close the file at end of the func
	readFile("./test.txt")
}

func checkErr(err error) {
	if err != nil {
		panic(err) //stop the program execution
	}
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkErr(err)
	fmt.Println("Content is :", string(databyte))
}
