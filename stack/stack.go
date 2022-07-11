package stack

import "errors"

type Stack struct {
	Arr []interface{}
	Cap int
}

func (s *Stack) IsEmpty() bool {
	return s.Cap == 0
}

func (s *Stack) Size() int {
	return s.Cap
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		panic(errors.New("stack is empty"))
	}
	return s.Arr[s.Cap-1]
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic(errors.New("stack is empty"))
	}
	s.Cap--
	ret := s.Arr[s.Cap]
	s.Arr = s.Arr[:s.Cap]
	return ret
}

func (s *Stack) Push(obj interface{}) {
	s.Arr = append(s.Arr, obj)
	s.Cap++
}
