// I want to understand better what range is doing.
package main

import (
	"fmt"
	"os"
)

func main(){
	for i, arg := range os.Args[1:] {
		fmt.Println(i)
		fmt.Println(arg)
	}
}
// ok so basically range in go = enumerate in python
// just realized that this program is also the anser to ex.1.2