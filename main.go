package main

import "fmt"

func variable()	string{
	// Decalre variable testing
	// Method 1, :=
	msg := "Drone CI/CD testing!!"
	// Method 2, var msg string
	var msg string
	msg = "Drone  CI/CD testing~~"
	return msg
}

func main() {
	fmt.Println(variable())
	fmt.Println("Running main.go successfully")
}