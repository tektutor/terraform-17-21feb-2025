package main

import "fmt"

func sayHello( msg string ) string {
	return "Hello, " + msg + " !";
}

func main() {
	fmt.Println( sayHello( "World" ) )
}
