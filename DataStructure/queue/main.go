package queue

import (
	"errors"
)

type Queue struct {
	elements []int
}

// Initailize
func Constructor() Queue {
	return Queue{}
}

func (this *Queue) Push(x int) {
	this.elements = append(this.elements, x)
}

func (this *Queue) Pop() (int, error) {
	if this.Empty() {
		return -1, errors.New("the queue is empty")
	}

	element := this.elements[0]

	if len(this.elements) > 1 {
		this.elements = this.elements[1:]
	} else {
		this.elements = this.elements[:0]
	}

	return element, nil
}

func (this *Queue) Peek() (int, error) {
	if !this.Empty() {
		return this.elements[0], nil
	}

	return -1, errors.New("the queue is empty")
}

func (this *Queue) Empty() bool {
	if len(this.elements) == 0 {
		return true
	}

	return false
}
