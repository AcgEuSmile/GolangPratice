package main

import (
	"fmt"
	"log"
)

func variable()	string{
	var msg string
	msg = "Drone  CI/CD testing~~"
	return msg
}

func variable2() string{
	msg := "Running main.go successfully"
	return msg
}

func assert(cond bool, msg string){
	_, f, l, _ := runtime.Caller(1)
	if !cond{
		fmt.Fprintf(os.Stderr, "Failed on (%s:%d): %s", f, l, msg)
	}
}

func main() {
	fmt.Println(variable())
	fmt.Println(variable2())
	assert(2*2==5, "the answer must be 4")
}