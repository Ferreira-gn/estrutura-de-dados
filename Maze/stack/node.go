package stack

type Position struct {
	X, Y int
}

type Node struct {
	value    Position
	nextNode *Node
}
