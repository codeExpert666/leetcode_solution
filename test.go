package main

import (
	"fmt"
	"github/DailyCode"
)

func main() {
	fmt.Println(DailyCode.CountGoodNodes([][]int{
		{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6},
	}))
}
