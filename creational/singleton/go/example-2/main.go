package main

import (
	"fmt"
	"time"
)

func main() {
	for range 30 {
		go getInstance()
		time.Sleep(time.Millisecond) // little delay to print the second one
	}

	fmt.Scanln()
}
