package main

import (
	"fmt"
	"rredis"
)

func main() {
	rredis.Init()
	fmt.Println("Hello world!")
}
