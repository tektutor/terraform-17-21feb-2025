package main

import (
  "fmt"
  "time"
)

func firstFunction( count int ) {
for i :=0; i<count; i++ {
     fmt.Println("First function", i)
     time.Sleep( time.Millisecond * 5 )
  }
}

func secondFunction() {
	for i := 0; i<10; i++ {
           fmt.Println ("Second Function")
	   time.Sleep( time.Millisecond * 5 )
	}
}

func main() {
   fmt.Println ("Press any key to exit ...")
   
   go firstFunction( 10 )
   go secondFunction()

   var tmp string
   fmt.Scanln(&tmp)
}
