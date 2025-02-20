package main

import "fmt"

//User-defined data type
type Rectangle struct {
   length int
   width  int
}

//This is a method that can be invoked with Rectangle instance variable
func (rect Rectangle) Area() int {
   area := rect.length * rect.width
   return area
}

//The below is a global function
func area( length, width int ) int {
   return length * width
}

func main() {

	rectangle := Rectangle {
		length: 100,
		width: 200,
	}

	fmt.Printf ("Area of rectangle: %d\n", rectangle.Area() )

	fmt.Printf ( "Area of some shape: %d\n", area( 10, 20 ) )
}
