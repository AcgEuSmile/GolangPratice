package main

import "fmt"

func variable()	string{
	var msg string
	msg = "Drone  CI/CD testing~~"
	return msg
}

func variable2() string{
	msg := "Running main.go successfully"
	return msg
}

func main() {
	fmt.Println(variable())
	fmt.Println(variable2())
}