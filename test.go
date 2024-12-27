package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool
	fmt.Printf("bool类型占用 %d 字节\n", unsafe.Sizeof(b))
}
