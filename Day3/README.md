# Day 3

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
