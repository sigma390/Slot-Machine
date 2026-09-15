package main

import (
	"fmt" //format package
)

func getName() string {
	name := "" // var name string
	fmt.Printf("Welcome to Om's casino\n")
	fmt.Printf("Enter Your name : ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Printf("%s \n", err)
		return ""
	}
	fmt.Printf("Welcome %s , lets play! \n", name)
	return name
}

func main() {
	var name string = getName()
	fmt.Printf("Hello %s\n", name)
}
