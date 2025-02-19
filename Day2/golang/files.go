package main

import (
  "fmt"
  "os"
  "log"
)


func main() {
	myfile, error := os.Create("./myfile.txt")

	if error != nil {
		log.Fatal(error)
	}

	//Declare a string and initialize with some text
	str := "This content will be stored in the file"

	bytesWritten, error := myfile.WriteString( str +"\n" )

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf( "Wrote %d bytes into the file\n", bytesWritten )
	myfile.Sync()
}
