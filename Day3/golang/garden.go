package main

import ( 
	"fmt"
	"time"
)

func digger(count int) {
	for i := 0; i<count ; i++ {
            fmt.Println ( "Digger digged pit ", i ) 
	    time.Sleep ( time.Millisecond * 3 )
	}
}

func planter(count int) {
	for i := 0; i< count; i++ {
            fmt.Println ( "Planter planted sapling", i ) 
	    time.Sleep ( time.Millisecond * 3 )
	}
}

func filler(count int) {
	for i := 0; i< count ; i++ {
            fmt.Println ( "Filler watered sapling", i ) 
	    time.Sleep ( time.Millisecond * 3 )
	}
}



func main() {
	go digger(500)
	go planter(500)
	go filler(500)
}
