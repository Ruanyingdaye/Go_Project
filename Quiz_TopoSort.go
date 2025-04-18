package main

import "fmt"

// 有向无环图的拓扑排序,还不算是入度出度的那个，一个节点的前面可能有X个parents

type Node struct {
	id      rune
	parents []rune
}

func topoSort(nodes []Node) {
	topoInMap := make(map[rune]int)
	//初始化每个node的时候，前置已经init了
	for _, node := range nodes {
		topoInMap[node.id] += len(node.parents)
	}
	// struct不能作为key，同样不能作为的还有function和map，channel都可以
	nodeMap := make(map[rune]Node)
	for _, node := range nodes {
		nodeMap[node.id] = node
	}

	for len(topoInMap) > 0 {
		for k, v := range topoInMap {
			fmt.Println(k, v)
			if v == 0 {
				fmt.Println(k)
				for _, parent := range nodeMap[k].parents {
					fmt.Println("parent", parent)
					topoInMap[parent]--
				}
				delete(topoInMap, k)
			}
		}
	}
}

func main() {
	nodeA := Node{'a', []rune{}}
	nodeB := Node{'b', []rune{'a'}}
	nodeC := Node{'c', []rune{'a'}}
	nodeD := Node{'d', []rune{'b', 'c'}}
	nodeE := Node{'e', []rune{'d'}}
	nodes := []Node{nodeA, nodeB, nodeC, nodeD, nodeE}
	topoSort(nodes)
}
