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

## Lab - Golang map
```
cd ~/terraform-17-21feb-2025
git pull
cd Day3/golang
go run ./map.go
```

Expected output
![image](https://github.com/user-attachments/assets/de89164a-b3dd-402b-bef2-c98e6dd69f34)

## Lab - Golang struct
```
cd ~/terraform-17-21feb-2025
git pull
cd Day3/golang
cat struct.go
go run ./struct.go
```

Expected output
![image](https://github.com/user-attachments/assets/c5c20c40-2f8b-4d2b-8dcd-6d62a492b2ba)

## Lab - Create custom go modules

In this exercise, we will be creating two custom go modules, namely hello and tektutor.  Hence in your home, create two folder at the same level namely hello and tektutor as shown below
```
cd ~
mkdir hello tektutor
cd hello
go mod init tektutor.org/hello
cat go.mod
go mod tidy
```

Now let's create a file named hello.go under hello folder as shown below
<pre>
package hello

import "fmt"

func SayHello( name string ) string {
   message := fmt.Sprintf("Hello, %v !", name )
   return message
}	
</pre>

Now, step out of hello folder and navigate to tektutor folder
```
cd ../tektutor
go mod init tektutor.org/tektutor
cat go.mod
go mod tidy
```

Let's create a main.go with the below code
<pre>
package main

import (
  "fmt"
  "tektutor.org/hello"
)

func main() {
	msg := hello.SayHello("Golang")
	fmt.Println(msg)
}	
</pre>

Now, let's run the below command under tektutor folder
```
cd ~/tektutor
go mod tidy
go mod edit --replace tektutor.org/hello=../hello
go mod tidy
go run ./main.go
```

Expected output
![image](https://github.com/user-attachments/assets/f397071b-edde-4b64-8f44-edc8ac411353)

## Lab - Module versioning
```
cd ~
cd hello
tree
cat v2/go.mod
cat v2/hello.go
cd ../tektutor
cat main.go
go mod edit --replace tektutor.org/hello/v2=../hello/v2
go mod tidy
go run ./main.go
```

Expected output
![image](https://github.com/user-attachments/assets/2d99e29a-267e-4f6e-8f03-dc5c20c43139)
![image](https://github.com/user-attachments/assets/6e2db257-db6d-422f-8588-aadfc7044924)
