package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web Request")

	url := "https://dummy.restapiexample.com/api/v1/employees"

	res, err := http.Get(url)
	checkErr(err)
	defer res.Body.Close() //Most imp responsibility to close the conx
	databytes, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	fmt.Println(string(databytes))
}

func checkErr(err error) {
	if err != nil {
		panic(err) //stop the program execution
	}
}
