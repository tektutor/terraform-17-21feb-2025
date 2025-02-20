package main

import "fmt"

func main() {
	i := 0

	for {
		if i == 5 {
			fmt.Println ( "Value of i is ", 5 )
			break
		} else {
		  fmt.Println ( i )
		  i++
		  continue
		}

	}
}
