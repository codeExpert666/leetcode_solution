package main

import (
	"fmt"
)

func main() {
	panicTest()
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
}

func panicTest() {
	panic("出错了")
}
