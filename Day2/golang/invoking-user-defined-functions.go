package main

import "fmt"

func main() {
	fmt.Println ( sayHello(" World") )
}

func sayHello(msg string) string {
	return "Hello, " + msg + " !"
}
