package pe

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Node struct {
	LeftIndex  int
	RightIndex int
	Min        int
	LeftNode   *Node
	RightNode  *Node
	Array      []int
}

func (n *Node) Print() {
	if n == nil {
		return
	}
	fmt.Printf("L: %d, R: %d, Min: %d, Arr: %v\n", n.LeftIndex, n.RightIndex, n.Min, n.Array)
	n.LeftNode.Print()
	n.RightNode.Print()
	return
}

func (n *Node) BuildTree() int {
	length := len(n.Array)
	if length == 1 {
		n.Min = n.Array[0]
		return n.Min
	}

	leftNode := &Node{
		LeftIndex:  n.LeftIndex,
		RightIndex: n.LeftIndex + length/2 - 1,
		Array:      n.Array[:length/2]}
	rightNode := &Node{
		LeftIndex:  n.LeftIndex + length/2,
		RightIndex: n.RightIndex,
		Array:      n.Array[length/2:]}
	n.LeftNode = leftNode
	n.RightNode = rightNode
	n.Min = min(leftNode.BuildTree(), rightNode.BuildTree())
	return n.Min
}

func (n *Node) GetRangeMin(left, right int) int {
	if left == n.LeftIndex && right == n.RightIndex {
		return n.Min
	}

	lr := n.LeftNode.RightIndex
	rl := n.RightNode.LeftIndex

	if left >= rl {
		return n.RightNode.GetRangeMin(left, right)
	}
	if right <= lr {
		return n.LeftNode.GetRangeMin(left, right)
	}
	return min(n.LeftNode.GetRangeMin(left, lr), n.RightNode.GetRangeMin(rl, right))
}

func NewSegmentTree(arr []int) *Node {
	n := &Node{
		LeftIndex:  0,
		RightIndex: len(arr) - 1,
		Array:      arr}
	n.BuildTree()
	return n
}
