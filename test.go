package main

import (
	"fmt"
	"sort"
)

func sortPeople(names []string, heights []int) []string {
	sort.Slice(names, func(i, j int) bool {
		return heights[i] > heights[j]
	})
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})
	fmt.Println(heights)
	return names
}

func main() {
	names := []string{"IEO", "Sgizfdfrims", "QTASHKQ", "Vk", "RPJOFYZUBFSIYp", "EPCFFt", "VOYGWWNCf", "WSpmqvb"}
	heights := []int{17233, 32521, 14087, 42738, 46669, 65662, 43204, 8224}
	fmt.Println(names)
	fmt.Println(heights)
	fmt.Println(sortPeople(names, heights))
}
