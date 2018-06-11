package main

import (
	"log"
)

const MaxInt = int(^uint(0) >> 1)

type MinStack struct {
	Data    []int
	MinData []int
	Min     int
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func Constructor() MinStack {
	// Invert bits (00...000 => 11...111), and shift left 1 bit (11...111 => 01...111)
	stack := MinStack{Min: MaxInt}

	return stack
}

func (this *MinStack) push(x int) {
	this.Data = append(this.Data, x)
	this.Min = Min(x, this.Min)
	this.MinData = append(this.MinData, this.Min)
}

func (this *MinStack) pop() {
	this.Data = this.Data[:len(this.Data)-1]
	this.MinData = this.MinData[:len(this.MinData)-1]

	if len(this.MinData) != 0 {
		this.Min = this.MinData[len(this.MinData)-1]
	} else {
		this.Min = MaxInt
	}
}

func (this *MinStack) peek() int {
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) getMin() int {
	return this.Min
}

func main() {
	stack := Constructor()

	stack.push(2)
	stack.push(2147483645)
	stack.push(2147483646)

	if stack.peek() != 2147483646 || stack.getMin() != 2 {
		log.Fatalf(
			"Expected peek[%v] != 2147483646 || min[%v] != 2",
			stack.peek(), stack.getMin(),
		)
	}

	stack.pop()

	if stack.peek() != 2147483645 || stack.getMin() != 2 {
		log.Fatalf(
			"Expected peek[%v] != 2147483645 || min[%v] != 2",
			stack.peek(), stack.getMin(),
		)
	}

	stack.pop()
	stack.pop()
	stack.push(2147483647)

	if stack.peek() != 2147483647 || stack.getMin() != 2147483647 {
		log.Fatalf(
			"Expected peek[%v] != 2147483647 || min[%v] != 2147483647",
			stack.peek(), stack.getMin(),
		)
	}

	stack.push(-2147483648)

	if stack.peek() != -2147483648 || stack.getMin() != -2147483648 {
		log.Fatalf(
			"Expected peek[%v] != -2147483648 || min[%v] != -2147483648",
			stack.peek(), stack.getMin(),
		)
	}

}
