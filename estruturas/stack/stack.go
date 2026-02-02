package structures_stack

import (
	"github.com/ferreira-gn/estrutura-de-dados/estruturas/structures"
)

type Stack[T any] struct {
	head *structures.Node[T]
	size int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		head: nil,
		size: 0,
	}
}

func (stack *Stack[T]) Push(value T) {
	newNode := &structures.Node[T]{
		Value: value,
		Next:  stack.head,
	}

	stack.head = newNode
	stack.size++
}

func (stack *Stack[T]) Pop() T {
	if stack.head == nil {
		var emptyValue T
		return emptyValue
	}

	returnedValue := stack.head.Value
	stack.head = stack.head.Next
	stack.size--

	return returnedValue
}

func (stack *Stack[T]) Peek() T {
	if stack.head == nil {
		var emptyValue T
		return emptyValue
	}

	return stack.head.Value
}

func (stack *Stack[T]) Len() int {
	return stack.size
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.size == 0
}
