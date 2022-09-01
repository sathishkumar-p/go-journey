package main

import "fmt"

//not allowed to write func inside the another func
//anonymous func is there
func main() {
	fmt.Println("Function")
	wel()
	res, msg := add(1, 2, 3, 4)
	fmt.Println(res, msg)
}
func wel() {
	fmt.Println("vanaakam ")
}

func add(values ...int) (int, string) {
	res := 0
	for _, v := range values {
		res = res + v
	}
	return res, "John"
}
