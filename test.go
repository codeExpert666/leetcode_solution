package main

import (
	"fmt"
	"github/Array"
)

func main() {
	a := []int{4, 2, 0, 2, 3, 2, 0}
	Array.NextPermutation([]int{4, 2, 0, 2, 3, 2, 0})
	fmt.Println(a)
}
