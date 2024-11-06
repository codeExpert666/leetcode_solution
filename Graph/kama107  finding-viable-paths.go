package main

import "fmt"

// 两个元素在同一集合中意味着两个元素之间存在“联系”，“联系”可抽象为一条边，从而形成图
// 并查集对“联系”进行了不同形式的表示，与其说并查集实在判断两个元素是否在同一集合中，
// 不如说并查集实在判断两个元素在图中是否连通
type UnionFindSet []int

// 切片是引用类型数据，其引用特性只体现在改变元素值会影响所有引用该值的切片
// 切片本质是结构体，结构体是值类型数据，所以要想使father在外部发生改变，必须使用指针接收
// 很好理解：虽然结构体是值类型，但结构体包含了数组首元素的地址，所以但改变其内部元素时会体现引用特性
// 当整体发生改变时，则体现为值类型
func (father *UnionFindSet) Init(length int) {
	*father = make(UnionFindSet, length)
	for i := 0; i < length; i++ {
		(*father)[i] = i
	}
}

func (father *UnionFindSet) find(node int) int {
	if node == (*father)[node] {
		return node
	}
	(*father)[node] = father.find((*father)[node]) // 压缩路径
	return (*father)[node]
}

func (father *UnionFindSet) Join(n1, n2 int) {
	n1 = father.find(n1)
	n2 = father.find(n2)
	(*father)[n1] = n2
}

func (father *UnionFindSet) IsConnected(n1, n2 int) bool {
	return father.find(n1) == father.find(n2)
}

func main() {
	var nodeNum, edgeNum int
	fmt.Scan(&nodeNum, &edgeNum)
	var father UnionFindSet
	father.Init(nodeNum + 1) // father[0]留空
	for i := 0; i < edgeNum; i++ {
		var n1, n2 int
		fmt.Scan(&n1, &n2)
		father.Join(n1, n2)
	}
	var source, destination int
	fmt.Scan(&source, &destination)
	if father.IsConnected(source, destination) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
