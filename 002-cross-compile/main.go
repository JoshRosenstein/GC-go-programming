package main

import (
	"fmt"
	"runtime"
)

func main() {
fmt.Printf("runtime: %sarchitecture: %s\n",runtime.GOOS,runtime.GOARCH)

}
