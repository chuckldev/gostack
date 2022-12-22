// Copyright (c) 2022 chuckldev.  All rights reserved.

package stack

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Name  string
	Value int
}

func TestPush(t *testing.T) {
	stk := New()

	if len(stk.Data) != 0 {
		t.Errorf("Stack should be zero ( 0 )")
	}

	stk.Push(TestStruct{Name: "number", Value: 1})

	if len(stk.Data) != 1 {
		t.Errorf("Stack should be one ( 1 )")
	}

	stk.Push(2)

	if len(stk.Data) != 2 {
		t.Errorf("Stack should be two ( 2 )")
	}
}

func TestPop(t *testing.T) {
	stk := New()
	wantName := "number"
	wantVal := 1

	if len(stk.Data) != 0 {
		t.Errorf("Stack should be zero ( 0 )")
	}

	stk.Push(5)

	if len(stk.Data) != 1 {
		t.Errorf("Stack should be one ( 1 )")
	}

	stk.Push(TestStruct{Name: "number", Value: 1})

	if len(stk.Data) != 2 {
		t.Errorf("Stack should be two ( 2 )")
	}

	val := stk.Pop().(TestStruct)

	if len(stk.Data) != 1 {
		t.Errorf("Stack should be one ( 1 )")
	}

	if val.Name != wantName {
		t.Errorf("val %s does not match %s", val.Name, wantName)
	}

	if val.Value != wantVal {
		t.Errorf("val %d does not match %d", val.Value, wantVal)
	}
}

func TestSize(t *testing.T) {
	stk := New()
	want := 2

	if len(stk.Data) != 0 {
		t.Errorf("Stack should be zero ( 0 )")
	}

	stk.Push(TestStruct{Name: "number", Value: 1})

	if len(stk.Data) != 1 {
		t.Errorf("Stack should be one ( 1 )")
	}

	stk.Push(2)

	if len(stk.Data) != 2 {
		t.Errorf("Stack should be two ( 2 )")
	}

	if stk.Size() != want {
		t.Errorf("Stack Size() should be two ( %d )", want)
	}
}

func TestIsEmpty(t *testing.T) {
	stk := New()
	want := true

	if stk.IsEmpty() != want {
		t.Errorf("Stack IsEmpty() should be %t", want)
	}

	stk.Push(1)
	want = false

	if stk.IsEmpty() != want {
		t.Errorf("Stack IsEmpty() should be %t", want)
	}
}

func TestPeek(t *testing.T) {
	stk := New()
	wantName := "number"
	wantVal := 1

	if len(stk.Data) != 0 {
		t.Errorf("Stack should be zero ( 0 )")
	}

	stk.Push(5)

	val := stk.Peek().(int)

	if val != 5 {
		t.Errorf("Stack Peek() should be five ( 5 )")
	}

	stk.Push(TestStruct{Name: "number", Value: 1})

	valPeek := stk.Peek().(TestStruct)

	if valPeek.Name != wantName {
		t.Errorf("val %s does not match %s", valPeek.Name, wantName)
	}

	if valPeek.Value != wantVal {
		t.Errorf("val %d does not match %d", valPeek.Value, wantVal)
	}

	stk.Pop()

	val = stk.Peek().(int)

	if val != 5 {
		t.Errorf("Stack Peek() should be five ( 5 )")
	}

}

func TestReverse(t *testing.T) {
	want := "4321"
	stk := New()

	stk.Push("1")
	stk.Push("2")
	stk.Push("3")
	stk.Push("4")

	res := stk.Reverse()
	var s string
	for !stk.IsEmpty() {
		s += fmt.Sprint(stk.Pop().(string))
	}

	if s != want {
		t.Errorf("Stack Reverse() failed.")
	}

	want = "1234"
	s = ""
	for !res.IsEmpty() {
		s += fmt.Sprint(res.Pop().(string))
	}

	if s != want {
		t.Errorf("Stack Reverse() failed.")
	}
}
