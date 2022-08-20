package main

import "fmt"

// global declaration is allowed but not allowed with valours operator
// token := test ---- no correct
// public variable has to start with capital letter

var GlobalValue int = 10

const TokenValue string = "hai"

func main() {

	//types and initialized
	var usr string = "sat"
	fmt.Println(usr)
	fmt.Printf("Variable Type %T \n", usr)

	var isVaule bool = true
	fmt.Println(isVaule)
	fmt.Printf("Variable Type %T \n", isVaule)

	var valueNum int = 256
	fmt.Println(valueNum)
	fmt.Printf("Variable Type %T \n", valueNum)

	var floatVal float32 = 123.789898098
	fmt.Println(floatVal)
	fmt.Printf("Variable Type %T \n", floatVal)

	// default values and not initialized it
	var val int
	fmt.Println(val)
	fmt.Printf("Variable Type %T \n", val)

	var isBool bool
	fmt.Println(isBool)
	fmt.Printf("Variable Type %T \n", isBool)

	var emptyStr string
	fmt.Println(emptyStr)
	fmt.Printf("Variable Type %T \n", emptyStr)

	var float32Empty float64
	fmt.Println(float32Empty)
	fmt.Printf("Variable Type %T \n", float32Empty)

	// implicit type
	// lexer take care of assigning the type of variable since it was initialized value.
	// lexer add semicolon as well

	var intValue = 2345
	fmt.Println(intValue)

	// no var style :=valours operator

	numValue := 1999
	fmt.Println(numValue)

	fmt.Println(TokenValue)
}
