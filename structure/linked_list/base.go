// 链表的基本实现

package main

import "fmt"

// LinkNode 链表结构体
type LinkNode struct {
	Data  int64
	NextNode *LinkNode
}

func main(){
	node := new(LinkNode)
	node.Data = 0

	// add new node
	node1 := new(LinkNode)
	node1.Data = 1
	node.NextNode  = node1

	node2 := new(LinkNode)
	node2.Data = 2
	node1.NextNode  = node2

	// 顺序打印
	for node != nil {
		fmt.Println(node.Data)
		node = node.NextNode
	}
}