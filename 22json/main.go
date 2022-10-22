package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSON")
	EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string `json:"course"` // alias for json obj
	Price    int
	Password string   `json:"-"`              // value will be removed in json convert
	Tags     []string `json:"tags,omitempty"` //if empty value then tags will be removed
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

func DecodeJson() {

	data := []byte(`
	{
		"course": "React",
		"Price": 200,
		"tags": ["web","sr"]
    }
	`)

	var courseData course

	checkValid := json.Valid(data)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(data, &courseData)
		//fmt.Println(courseData) //print value
		fmt.Printf("%#v\n", courseData) //print additional information about the struct
	} else {
		fmt.Println("JSON was not valid")
	}
	//some cases where you just want to add data to key value
	var myData map[string]interface{}
	json.Unmarshal(data, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key is %v and value is %v and Type is: %T", k, v, v)
	}
}
