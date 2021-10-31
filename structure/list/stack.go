package list

import "errors"

type Stack struct {
	elements []interface{}
	length   int
}

func NewStack(elements ...interface{}) *Stack {
	s := &Stack{nil, 0}
	for _, data := range elements {
		s.Push(data)
	}
	return s
}

func (s *Stack) Push(data interface{}) {
	s.elements = append(s.elements, data)
	s.length++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("下溢操作")
	}
	data := s.elements[s.length-1]
	s.elements = s.elements[:s.length-1]
	s.length--
	return data, nil
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack) Clear() {
	s.length = 0
	s.elements = nil
}
