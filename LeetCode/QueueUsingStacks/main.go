package main

import (
	"fmt"

	"github.com/ferreira-gn/estrutura-de-dados/LeetCode/QueueUsingStacks/myqueue"
)

func main() {
	q := myqueue.Constructor()

	fmt.Println("Empty:", q.Empty())

	q.Push(1)
	q.Push(2)

	fmt.Println("Peek:", q.Peek())
	fmt.Println("Pop:", q.Pop())
	fmt.Println("Empty:", q.Empty())

	fmt.Println("Pop:", q.Pop())
	fmt.Println("Empty:", q.Empty())
}
