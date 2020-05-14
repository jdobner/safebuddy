package main

import (
	"fmt"

	"github.com/golang/glog"
)

func main() {
	SayHello()
}

func SayHello() {
	a := fmt.Sprintf("jj %v JJ", "hh")
	fmt.Println(a)
	glog.Error("Testing")
}
