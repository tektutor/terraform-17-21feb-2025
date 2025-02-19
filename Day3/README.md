# Day 3

## Info - Golang basic data types
<pre>
- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64,, uintptr
- byte ( alias of uint8 )
- rune ( alias for int32 - represents a unicode )
- float32, float64
- complex64, complex128
</pre>

## Lab - Writing an user-defined function and invoking it from main in golang

Create a file named invoking-user-defined-functions.go with the below content
```
package main

import "fmt"

func main() {
	fmt.Println ( sayHello(" World") )
}

func sayHello(msg string) string {
	return "Hello, " + msg + " !"
}
```

Run the program
```
go run ./invoking-user-defined-functions.go
```

Expected output
![image](https://github.com/user-attachments/assets/5652a04e-1691-46aa-9463-ce6b979acb9d)


## Lab - Loops in golang
```
cd ~/terraform-17-21feb-2025
git pull
cd Day3/golan
cat ./loops.go
go run ./loops.go
```

Expected output
![image](https://github.com/user-attachments/assets/37d2e280-6879-45eb-8919-459bf8ecae54)
