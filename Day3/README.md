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


## Lab - golang Loops
```
cd ~/terraform-17-21feb-2025
git pull
cd Day3/golan
cat ./loops.go
go run ./loops.go
```

Expected output
![image](https://github.com/user-attachments/assets/37d2e280-6879-45eb-8919-459bf8ecae54)

## Lab - Golang Switch case
```
cd ~/terraform-17-21feb-2025
git pull
cd Day3/golang
cat switch-case.go
go run ./cat switch-case.go 
```

Expected output
![image](https://github.com/user-attachments/assets/4908f4ff-e66d-4d02-aa74-35edecd1bc4b)

## Lab - Golang arrays
```
cd ~/terraform-17-21feb-2025
git pull
cd Day3/golang
cat arrays.go
go run ./arrays.go
```

Expected output
![image](https://github.com/user-attachments/assets/b79deb0b-09c5-4945-9270-281bb96995ee)
![image](https://github.com/user-attachments/assets/1907bb33-cc15-4413-831c-7b057c47f18c)

## Lab - Golang slice
```
cd ~/terraform-17-21feb-2025
git pull
cd Day3/golang
cat slice.go
go run ./slice.go
```

Expected output
![image](https://github.com/user-attachments/assets/c5d51db8-5278-4608-bded-c747835ba3fe)
![image](https://github.com/user-attachments/assets/c0290c81-49c2-4d3b-b074-f6da5083a60d)
![image](https://github.com/user-attachments/assets/34879507-862e-4f77-9e3e-c0a9b1b1033a)
