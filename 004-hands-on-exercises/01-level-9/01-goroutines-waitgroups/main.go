package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
fmt.Println("Begin CPU:",runtime.NumCPU())
fmt.Println("Begin gt:",runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello From Thing one ")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from Thing two")
		wg.Done()
	}()

	fmt.Println("about to exit")
	fmt.Println("Mid CPU:",runtime.NumCPU())
	fmt.Println("Mid gt:",runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End CPU:",runtime.NumCPU())
	fmt.Println("End gt:",runtime.NumGoroutine())
}