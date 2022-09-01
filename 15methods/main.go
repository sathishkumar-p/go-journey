package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent
	sathish := User{"Sathish", "sathish@devops.com", true, 24}

	fmt.Println(sathish)
	fmt.Printf("More Info: %+v\n", sathish) // +v gives more info like include the variable name
	fmt.Printf("More Info: %v\n", sathish)
	sathish.GetStatus()
	sathish.SetMail()
	fmt.Printf("More Info: %v\n", sathish)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//methods if the func uses the struct then its called as methods
func (u User) GetStatus() {
	fmt.Println(u.Status)
}

func (u User) SetMail() {
	u.Email = "test@gmail.com"
	fmt.Println(u.Email)
}
