package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSON")
	EncodeJson()
}

type course struct {
	Name     string `json:"course"` // alias for json obj
	Price    int
	Password string   `json:"-"` // value will be removed in json convert
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	courses := []course{
		{"React", 200, "acs", []string{"web", "sr"}},
		{"GO", 200, "acs", nil},
	}

	// jsonVal, _ := json.Marshal(courses)  convert as json not indent
	jsonVal, _ := json.MarshalIndent(courses, "", "\t")

	fmt.Printf("%s\n", jsonVal)
}
