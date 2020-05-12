package main

import "fmt"

func deferAndReturn() (int,error) {
	defer fmt.Println("print1")
	defer fmt.Println("Print2")
	return fmt.Println("Return")
}

func main() {
	deferAndReturn()
	fmt.Println("Return is ahead of defer")
	fmt.Println("defer is FILO")
}
