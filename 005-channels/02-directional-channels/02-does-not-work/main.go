package main

import "fmt"

func main() {
	c := make(chan<- int, 2)  /// This channel is a send only channel

	c <- 42
	c <- 43

	fmt.Println(<-c) /// doesnt work because channel could only send
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}