package main

import (
	"fmt"

	"bookProjects/ch2/popcount/pop"
)


func main() {
	var pc [256]byte
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	fmt.Println(pc[170])
	fmt.Println(pop.PopCount(uint64(170)))
	
}
