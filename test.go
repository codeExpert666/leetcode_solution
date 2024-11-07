package main

import (
	"fmt"
	"github/ArrayList"
)

func main() {
	n4 := ArrayList.ListNode{Val: 1, Next: nil}
	n3 := ArrayList.ListNode{Val: 2, Next: &n4}
	n2 := ArrayList.ListNode{Val: 2, Next: &n3}
	n1 := ArrayList.ListNode{Val: 1, Next: &n2}
	fmt.Println(ArrayList.IsPalindrome(&n1))
}
