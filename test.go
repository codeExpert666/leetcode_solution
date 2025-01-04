package main

import (
	"fmt"
	"github/Array"
)

func main() {
	mc := Array.ConstructorMCT()
	fmt.Println(mc.Book(24, 40))
	fmt.Println(mc.Book(43, 50))
	fmt.Println(mc.Book(27, 43))
	fmt.Println(mc.Book(5, 21))
	fmt.Println(mc.Book(30, 40))
	fmt.Println(mc.Book(14, 29))
	fmt.Println(mc.Book(3, 19))
}
