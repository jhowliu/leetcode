package test

import (
	"testing"

	"./queue"
	"./stack"
)

func TestQueue(t *testing.T) {
	q := queue.Constructor()

	q.Push(10)
	q.Push(10000)

	e, err := q.Pop()

	if err != nil {
		t.Fatalf("Pop occur error: %v", err)
	}

	if e != 10 {
		t.Fatalf("Pop: Element:%v is not equal to 10", e)
	}

	if q.Empty() != false {
		t.Fatalf("The queue should not be empty: %v", q.Empty())
	}

	e, err = q.Peek()

	if err != nil {
		t.Fatalf("Peek occur error: %v", err)
	}

	if e != 10000 {
		t.Fatalf("Peek: Element:%v is not equal to 10000", e)
	}

	e, err = q.Pop()

	if err != nil {
		t.Fatalf("Pop occur error: %v", err)
	}

	if e != 10000 {
		t.Fatalf("Pop: Element:%v is not equal to 10", e)
	}

	if q.Empty() != true {
		t.Fatalf("The queue should be empty: %v", q.Empty())
	}
}

func TestStack(t *testing.T) {
	s := stack.Constructor()

	s.Push(10)
	s.Push(10000)

	e := s.Pop()

	if e != 10000 {
		t.Fatalf("Pop: Element:%v is not esual to 10000", e)
	}

	if s.Empty() != false {
		t.Fatalf("The sueue should not be empty: %v", s.Empty())
	}

	e = s.Peek()

	if e != 10 {
		t.Fatalf("Peek: Element:%v is not esual to 10", e)
	}

	e = s.Pop()

	if e != 10 {
		t.Fatalf("Pop: Element:%v is not esual to 10", e)
	}

	if s.Empty() != true {
		t.Fatalf("The stack should be empty: %v", s.Empty())
	}
}
