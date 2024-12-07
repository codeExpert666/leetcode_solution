package main

import (
	"fmt"
	"github/StringFunc"
)

func main() {
	fmt.Println(StringFunc.GetFolderNames([]string{"kaido", "kaido(1)", "kaido", "kaido(1)", "kaido(2)"}))
}
