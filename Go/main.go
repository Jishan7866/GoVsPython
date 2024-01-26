package main

import (
	"fmt"
	"time"
)

func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += 1
	}
	return total
}

func main() {
	n := 1000000000
	st := time.Now()
	total := sum(n)
	en := time.Now()
	fmt.Printf("Time Take With Go: %.2f secconds \n", en.Sub(st).Seconds())
	fmt.Println("Go Result: ", total)
}
