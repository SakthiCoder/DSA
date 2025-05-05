package stack

type Stack struct {
	items []any
}

func (s *Stack) Push(num any) {
	s.items = append(s.items, num)
}

func (s *Stack) Pop() (any, bool) {

	if len(s.items) == 0 {
		return -1, false
	}
	item := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return item, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peek() any {

	if len(s.items) == 0 {
		return -1
	}

	return s.items[len(s.items)-1]
}
