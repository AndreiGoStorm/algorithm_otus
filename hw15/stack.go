package hw15

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	index := len(s.data) - 1
	elem := s.data[index]
	s.data = s.data[:len(s.data)-1]

	return elem
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
