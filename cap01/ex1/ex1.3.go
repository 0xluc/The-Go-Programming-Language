package main

import (
	"fmt"
	"time"
	"os"
	"strings"
)

func echo1(){
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
func echo3(){
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	startEcho1 := time.Now()

	echo1()

	durationEcho1 := time.Since(startEcho1)
	fmt.Println("durationEcho1:", durationEcho1.Microseconds())
	startEcho3 := time.Now()

	echo3()

	durationEcho3 := time.Since(startEcho3)
	fmt.Println("durationEcho3:", durationEcho3.Microseconds())
}
// so basically the second one is 10x faster