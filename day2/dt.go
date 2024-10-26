package main

import (
	"fmt"
	"reflect"
)

func main() {
	var expStr string = "Hello"
	impStr := "Hello"

	const maxVal = 100

	fmt.Printf("explicitString: Type=%v, Value=%v\n", reflect.TypeOf(expStr), expStr)
	fmt.Printf("implicitString: Type=%v, Value=%v\n", reflect.TypeOf(impStr), impStr)
	fmt.Printf("maxValue: Type=%v, Value=%v\n", reflect.TypeOf(maxVal), maxVal)

	var integer int = 42
	var float float32 = float32(integer)
	var backToInt int = int(float)

	fmt.Printf("Conversion chain: %v (int) -> %v (float32) -> %v (int)\n", integer, float, backToInt)

	fuck_ios()
	error_file();
}
