// defer keyword  stat will exe end of the function or later.
// it will create macdonald queue. Follows LIFO strategy
package main

import "fmt"

func main() {
	fmt.Println("Defer")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	deferfunc()
}

func deferfunc() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
