package stack

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(num interface{}) {
	s.items = append(s.items, num)
}

func (s *Stack) Pop() (interface{}, bool) {

	if len(s.items) == 0 {
		return -1, false
	}
	item := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return item, true
}

func (s *Stack) IsEmpty() interface{} {
	return len(s.items) == 0
}

func (s *Stack) Peek() interface{} {

	if len(s.items) == 0 {
		return -1
	}

	return s.items[len(s.items)-1]
}
