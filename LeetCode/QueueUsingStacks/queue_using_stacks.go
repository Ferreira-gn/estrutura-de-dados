package myqueue

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) Pop() int {
	q.moveIfNeeded()

	top := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]

	return top
}

func (q *MyQueue) Peek() int {
	q.moveIfNeeded()
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

func (q *MyQueue) moveIfNeeded() {
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			n := len(q.inStack) - 1
			value := q.inStack[n]
			q.inStack = q.inStack[:n]
			q.outStack = append(q.outStack, value)
		}
	}
}
