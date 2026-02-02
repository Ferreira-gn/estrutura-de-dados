package structures_linkedlist

import (
	"github.com/ferreira-gn/estrutura-de-dados/estruturas/structures"
)

type LinkedList[T any] struct {
	head *structures.Node[T]
	end  *structures.Node[T]
	size int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		end:  nil,
		size: 0,
	}
}

func (linkedList *LinkedList[T]) Push(value T) {
	newNode := &structures.Node[T]{
		Value: value,
		Next:  nil,
	}

	if linkedList.head == nil {
		linkedList.head = newNode
		linkedList.end = newNode
		linkedList.size = 1
		return
	}

	if linkedList.end == linkedList.head {
		linkedList.head.Next = newNode
		linkedList.end = newNode
		linkedList.size++
		return
	}

	linkedList.end.Next = newNode
	linkedList.end = newNode
	linkedList.size++
}

func (linkedList *LinkedList[T]) Pop() {
	if linkedList.head == nil {
		return
	}

	if linkedList.end == linkedList.head {
		linkedList.end = nil
		linkedList.head = nil
		return
	}

	currentNode := linkedList.head

	for i := 0; currentNode.Next != nil; i++ {

		if currentNode.Next.Next == nil {
			currentNode.Next = nil
			break
		}

		currentNode = currentNode.Next
	}

	linkedList.end = currentNode
}

func (linkedList *LinkedList[T]) Len() int {
	if linkedList.head == nil {
		return 0
	}

	return linkedList.size
}

func (linkedList *LinkedList[T]) GetElementByIndex(index int) T {
	if linkedList.head == nil {
		var emptyElement T
		return emptyElement
	}

	currentIndex := 0
	currentNode := linkedList.head

	for currentIndex != index {
		currentNode = currentNode.Next
		currentIndex++
	}

	return currentNode.Value
}
