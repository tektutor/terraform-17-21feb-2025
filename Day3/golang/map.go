package main

import "fmt"


func main() {

	toolsPath := map[string]string {
		"java_home": "/usr/lib/jdk-11",
		"mvn_home" : "/usr/share/maven",
	}

	//Given a key, map will be able to retrieve the value
	fmt.Println ( "Java home directory :", toolsPath["java_home"] )

	//add a key,value pair into an existing map
	toolsPath["go_home"] = "/usr/go"

	//iterating a map and printing its values
	for key,value := range toolsPath {
		fmt.Println ( key, value )
	}

	//delete a key-value pair
	delete ( toolsPath, "go_home")
	fmt.Println (toolsPath)
}
