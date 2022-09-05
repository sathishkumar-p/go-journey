package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("web request call")
	GetAll()
	Post()
}

func GetAll() {
	const url = "https://reqres.in/api/users?page=2"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status :", res.StatusCode)

	var resString strings.Builder
	con, _ := ioutil.ReadAll(res.Body)
	byteCount, _ := resString.Write(con)

	fmt.Println(byteCount)
	fmt.Println(resString.String())
}

func Post() {
	const url = "https://reqres.in/api/users"
	//json payload

	body := strings.NewReader(`
		{
			"name": "morpheus",
			"job": "leader"
		}
	`)

	res, _ := http.Post(url, "application/json", body)

	defer res.Body.Close()

	databytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(databytes))
}

func PostForm() {
	const urls = "https://reqres.in/api/users"
	//formdata

	data := url.Values{}
	data.Add("fn", "sat")
	data.Add("ln", "sat")

	res, _ := http.PostForm(urls, data)

	fmt.Println(res.StatusCode)
}
