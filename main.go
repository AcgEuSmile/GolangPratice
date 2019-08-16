package main

import "fmt"

func variable()	string{
	msg := "Drone CI/CD testing!!"
	return msg
}

func main() {

	fmt.Println(variable())
	fmt.Println("Running main.go successfully")
}