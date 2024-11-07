package main

import "fmt"

type UnionFindSet []int

func (s *UnionFindSet) Init(length int) {
	*s = make(UnionFindSet, length)
	for i := 0; i < length; i++ {
		(*s)[i] = i
	}
}

func (s *UnionFindSet) find(n int) int {
	if n == (*s)[n] {
		return n
	}
	(*s)[n] = s.find((*s)[n])
	return (*s)[n]
}

func (s *UnionFindSet) Join(n1, n2 int) {
	n1 = s.find(n1)
	n2 = s.find(n2)
	(*s)[n1] = n2
}

func (s *UnionFindSet) IsConnected(n1, n2 int) bool {
	return s.find(n1) == s.find(n2)
}

func main() {
	var nodeNum int
	fmt.Scan(&nodeNum)
	var s UnionFindSet
	s.Init(nodeNum + 1)
	var ansNode1, ansNode2 int
	for i := 0; i < nodeNum; i++ {
		var n1, n2 int
		fmt.Scan(&n1, &n2)
		if s.IsConnected(n1, n2) { // 若输入边的两个顶点已经连通，则说明该边是冗余边
			ansNode1, ansNode2 = n1, n2
		} else {
			s.Join(n1, n2)
		}
	}
	fmt.Printf("%d %d", ansNode1, ansNode2)
}
