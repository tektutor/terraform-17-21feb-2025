package main

import "fmt"

func main() {
	var inventor1 = "Ken Thompson" //Declares a string implicitly

	inventor2 := "Rob Pike"	       //Short form of declaring a variable and initializing it with some value
	var inventor3 string	       //Declares string variable
	inventor3 = "Robert Griesemer" //initialize the string variable with a value

	fmt.Println ("Golang inventors")
	fmt.Println ( inventor1 )
	fmt.Println ( inventor2 )
	fmt.Println ( inventor3 )

	var firstNumber, secondNumber int = 10, 20
	fmt.Println ( "Value of first number is ", firstNumber )
	fmt.Println ( "Value of second number is ", secondNumber )

	str1 := "Hello,"
	str2 := "World"

	isSuccess := true

	fmt.Println ( str1 + " " + str2 )

	//Prints the values
	fmt.Printf("Inventor1 = %v, Inventor2 = %v, Inventor3 = %v\n", inventor1, inventor2, inventor3 )

	//Prints the variable types 
	fmt.Printf("Inventor1 = %T, Inventor2 = %T, Inventor3 = %T\n", inventor1, inventor2, inventor3 )
	fmt.Printf("%T\n", firstNumber )
	fmt.Printf("%T\n", secondNumber )
	fmt.Printf("%T\n", isSuccess )


}
