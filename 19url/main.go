package main

import (
	"fmt"
	"net/url"
)

const urlValue string = "https://reqres.in/api/users?page=2"

func main() {
	fmt.Println("Url Parse")

	fmt.Println(urlValue)

	//parsing
	res, _ := url.Parse(urlValue)

	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Port())
	fmt.Println(res.Path)
	fmt.Println(res.RawQuery)

	qparams := res.Query()
	fmt.Printf("Type of the query %T\n", qparams)

	fmt.Println(qparams["page"])

	for _, v := range qparams {
		fmt.Println(v)
	}
	// construct  url
	// remember &
	ownUrl := &url.URL{
		Scheme:  "https",
		Host:    "reqres.in",
		Path:    "/api",
		RawPath: "user=sat",
	}
	fmt.Println(ownUrl.String())
}
