package main

import (
	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 3,
				Right: &Node{
					Val: 4,
				},
			},
		},
	}

	spew.Dump(getLeftNodes(root))
}

func getLeftNodes(root *Node) []*Node {
	if root == nil {
		return nil
	}

	leftNodes := []*Node{}
	bfs(root, 0, &leftNodes)

	return leftNodes
}

func bfs(node *Node, level int, leftNodes *[]*Node) {
	if len(*leftNodes) == level {
		*leftNodes = append(*leftNodes, node)
	}

	if node.Left != nil {
		bfs(node.Left, level+1, leftNodes)
	}

	if node.Right != nil {
		bfs(node.Right, level+1, leftNodes)
	}
}
