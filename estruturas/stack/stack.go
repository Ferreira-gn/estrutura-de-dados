package stack

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	head *Node
	size int
}

func StackConstructor() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Push(value int) {
	newNode := Node{value, nil}

	if s.head == nil {
		s.head = &newNode
		s.size++
		return
	}

	newNode.next = s.head
	s.head = &newNode
	s.size++
}

func (s *Stack) Pop() {
	if s.head == nil {
		return
	}

	s.head = s.head.next
	s.size--
}

func (s *Stack) ViewTheTop() int {
	if s.head == nil {
		return 0
	}

	return s.head.value
}

func (s *Stack) Len() int {
	return s.size
}
