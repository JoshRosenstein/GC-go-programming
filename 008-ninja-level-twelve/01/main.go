package main

import (
	"fmt"
	"github.com/WaitNKill2/go-programming/008-ninja-level-twelve/01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
