package structures_queue

import (
	"github.com/ferreira-gn/estrutura-de-dados/estruturas/structures"
)

type Queue[T any] struct {
	head *structures.Node[T]
	end  *structures.Node[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		end:  nil,
	}
}

func (queue *Queue[T]) Push(value T) {
	newNode := &structures.Node[T]{
		Value: value,
		Next:  nil,
	}

	if queue.head == nil {
		queue.head = newNode
		queue.end = newNode
		return
	}

	if queue.end == queue.head {
		queue.head.Next = newNode
		queue.end = newNode
		return
	}

	queue.end.Next = newNode
	queue.end = newNode
}

func (queue *Queue[T]) Dequeue() T {
	if queue.head == nil {
		var emptyValue T
		return emptyValue
	}

	if queue.end == queue.head {
		returnedValue := queue.head.Value
		queue.head = nil
		queue.end = nil

		return returnedValue
	}

	returnedValue := queue.head.Value
	queue.head = queue.head.Next
	return returnedValue
}
