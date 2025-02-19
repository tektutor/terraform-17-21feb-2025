package main

import "fmt"

func main() {
	
	//shorter syntax to declare and initialize an array of int with 5 values
	intArray := [6] int { 10, 20, 30, 40, 50, 60 }
/*
	var anotherArray [3]int
	fmt.Println ( "another array elements are ..." )
	fmt.Println ( anotherArray )
*/

	fmt.Println ( "Array elements are .." )
	fmt.Println ( intArray )

	//Slice uses array internally
	//in this case, the slice is referring to an array from index position 2 to 4
	var mySlice []int = intArray[2:5]
	fmt.Println ( "Slice elements are ... ")
	fmt.Println ( mySlice )

	mySlice[0] = 100
	mySlice[1] = 200
	mySlice[2] = 300
	//mySlice[3] = 400 - give index out of bounds error
	//mySlice[4] = 500 - give index out of bounds error


	mySlice = append( mySlice, 400, 500 )

	fmt.Println ( "Array elements are .." )
	fmt.Println ( intArray )
	fmt.Println ( "Slice elements are ... ")
	fmt.Println ( mySlice )

	mySlice[0] = 0
	mySlice[1] = 0
	mySlice[2] = 0
	fmt.Println ( "Array elements are .." )
	fmt.Println ( intArray )
	fmt.Println ( "Slice elements are ... ")
	fmt.Println ( mySlice )

/*
	//golang slice can dynamically grow in size
	anotherSlice := [] int { 1, 2, 3 }
	fmt.Println ( "Another slice elements ...")
	fmt.Println ( anotherSlice )

	//append some additional values into the slice
	anotherSlice = append( anotherSlice, 4, 5, 6 )
	fmt.Println ( "Another slice elements after appending values are ...")
	fmt.Println ( anotherSlice )
*/

}
