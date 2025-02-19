package main
import "fmt"

func main() {
	str1 := "go"
	str2 := "go"
	fmt.Printf ( "%v and %v matching? : %v\n", str1, str2, compareString ( str1, str2 ) )

	str1 = "Golang"
	str2 = "go"
	fmt.Printf ( "%s and %s matching? : %v\n", str1, str2, compareString ( str1, str2 ) )
}

func compareString( str1, str2 string ) bool {
	//We are declaring an user-defined bool variable named areSame
	var areSame bool

	if str1 == str2 {
		areSame = true
	} else {
		areSame = false
	}

	return areSame
}

