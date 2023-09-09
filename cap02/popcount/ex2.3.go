package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func main() {
	start1 := time.Now()
	PopCount(10)
	duration1 := time.Since(start1)
	fmt.Println("duration1:", duration1)
	start2 := time.Now()
	PopCount2(10)
	duration2 := time.Since(start2)
	fmt.Println("duration2:", duration2)
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	count := 0
	for i := 0; i <= 7; i++ {
		count += int(pc[byte(x>>(uint(i)))])
	}
	return count
}

func PopCount2(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}
