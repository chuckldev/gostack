// Copyright (c) 2022 chuckldev.  All rights reserved.

package stack

type Stack struct {
	Data []interface{}
}

func New() *Stack {
	return &Stack{nil}
}

func (s *Stack) Push(data interface{}) {
	s.Data = append(s.Data, data)
}

func (s *Stack) Pop() interface{} {
	res := s.Data[len(s.Data)-1 : len(s.Data)]
	s.Data = s.Data[:len(s.Data)-1]
	return res[0]
}

func (s *Stack) Size() int {
	return len(s.Data)
}

func (s *Stack) IsEmpty() bool {
	if len(s.Data) == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Reverse() Stack {
	temp := New()
	for !s.IsEmpty() {
		temp.Push(s.Pop())
	}
	return *temp
}

func (s *Stack) Peek() interface{} {
	res := s.Data[len(s.Data)-1 : len(s.Data)]
	return res[0]
}
