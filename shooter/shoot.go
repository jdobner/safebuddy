package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	sayHello()
}

func sayHello() {
	a := fmt.Sprintf("jj %v JJ", "hh")
	fmt.Println(a)
	b := proto.Float32(22)
	fmt.Println(b)
	sayhi()

}
