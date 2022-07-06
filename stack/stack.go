package stack

type Stack struct {
	Arr []interface{}
	Cap int
}

func (s *Stack) Pop() interface{} {
	if s.Cap == 0 {
		return nil
	}
	s.Cap--
	return s.Arr[s.Cap]
}

func (s *Stack) Push(obj interface{}) {
	s.Arr = append(s.Arr, obj)
	s.Cap++
}
