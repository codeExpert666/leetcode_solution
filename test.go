package main

import (
	"fmt"
	// "github/ArrayList"
)

func main() {
	a1 := [][2]int{}
	a2 := [][2]int{{1, 2}, {3, 4}}
	a1 = append(a1, a2[0])
	fmt.Println(a1)
}
