package main

import (
	"fmt"
	"github/DailyCode"
)

func main() {
	fmt.Println(DailyCode.ImageSmoother([][]int{
		{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
	}))
}
