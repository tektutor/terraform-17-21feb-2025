package main

import "fmt"

func main() {

	// -128 to 127
//	var x int8

	// Range is -32768 to 32767
//	var y int16

	//-2147483648 to 2147483647
//	var z int32

	//In a 32-bit OS int is considered as int32
	//In a 64-bit OS int is considered as int64
//	var i int

	var a, b, result int16

	a = 100
	b = 100
	result = a + b

	fmt.Printf("Value of a is %d\n", a )
	fmt.Printf("Value of b is %d\n", b )
	fmt.Printf("sum of %d and %d is %d\n", a, b, result )

	//Range is 0 to 255
	var z uint8

	z = 256 //not supported

	fmt.Printf("Value of %d\n", z )

}
