package main

import "fmt"

func main() {
	SayHello()
}

func SayHello() {
	a := fmt.Sprintf("jj %v JJ", "hh")
	fmt.Println(a)
}
