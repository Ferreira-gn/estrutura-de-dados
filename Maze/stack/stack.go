package stack

import "fmt"

type Stack struct {
	head *Node
	end  *Node
	size int
}

func NewStack(initialPosition Position) *Stack {
	node := &Node{value: initialPosition}

	return &Stack{
		head: node,
		end:  node,
		size: 1,
	}
}

func (stack Stack) Len() int {
	return stack.size
}

func (stack Stack) ViewTheTop() Position {
	return stack.end.value
}

func (stack *Stack) Append(x, y int) {
	newNode := &Node{
		value: Position{x, y},
	}

	if stack.head == nil || stack.end == nil {
		stack.head = newNode
		stack.end = newNode
		stack.size++
		return
	}

	stack.end.nextNode = newNode
	stack.end = newNode
}

func (stack *Stack) Pop() {
	if stack.head == nil {
		fmt.Printf("A lista está vazia")
		return
	}

	currentNode := stack.head

	for currentNode != nil {

		if currentNode.nextNode == stack.end {
			currentNode.nextNode = nil
			stack.end = currentNode
			return
		}

		currentNode = currentNode.nextNode
	}

}
