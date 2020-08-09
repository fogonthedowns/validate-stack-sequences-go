package main

func validateStackSequences(pushed []int, popped []int) bool {
	stack := New()
	var count int

	for _, value := range pushed {
		stack.Push(&Item{value: value})
		for {
			if stack.count == 0 || count > len(pushed) || stack.Peek().value != popped[count] {
				break
			}
			stack.Pop()
			count++
		}
	}
	return count == len(pushed)
}

// Stack

type Item struct {
	value int
	next  *Item
}

type Stack struct {
	first *Item
	count int
}

func New() (s *Stack) {
	return &Stack{}
}

func (s *Stack) Push(i *Item) {
	if s.count == 0 {
		s.first = i
		s.count = 1
		return
	}

	previous := s.first
	s.first = i
	s.first.next = previous
	s.count++
	return
}

func (s *Stack) Pop() (i *Item) {
	if s.count == 0 {
		return nil
	}
	s.count--
	previous := s.first
	s.first = s.first.next
	return previous
}

func (s *Stack) Peek() (i *Item) {
	if s.count == 0 {
		return nil
	}

	return s.first
}
