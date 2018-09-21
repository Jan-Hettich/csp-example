package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int, 10)
	sequence := make(chan int)

	// Generate sequence
	go func() {
		for x := 1; x <= 100; x++ {
			fmt.Println("sending", x)
			naturals <- x
			// time.Sleep(100 * time.Millisecond)
		}
		close(naturals)
	}()

	// Copy to new sequence
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			// time.Sleep(200 * time.Millisecond)
			sequence <- x
		}
		close(sequence)
	}()

	// Print
	for x := range sequence {
		fmt.Println("receiving", x)
	}

}
