package main

import "fmt"

func main() {

	count := 5  // We are declaring a variable count and assigning a integer value 5

	for count > 0 {

		fmt.Println ( count )
		count-- // equivalent to count = count - 1
	}

	fmt.Println("Value of count is ", count, " after for loop ")

	fmt.Println("-----------------------------------------------------------------")

	count = 0 // count variable is aready above, we are just reinitializing to value 0

	for count = 1; count < 10; count++ {
		fmt.Printf("%d\t", count )
	}

	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
}
