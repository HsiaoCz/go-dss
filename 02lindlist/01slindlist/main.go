package main

import "fmt"

// 单链表
// 单链表，一个节点保存数据以及下一个节点的引用

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	// 新的节点
	node := new(LinkNode)
	node.Data = 2

	// 新的节点
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1

	// 新的节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2

	// 顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			// 打印系欸但值
			fmt.Println(nowNode.Data)
			// 获取下一个节点
			nowNode = nowNode.NextNode
			continue
		}
		// 如果下一个节点为空
		break
	}
}
