package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		time.Sleep(time.Second)
		fmt.Printf("My favorite number is %d\n", i)
		i++
	}
}
