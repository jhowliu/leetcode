package stack

type Stack struct {
	elements []int
}

func Constructor() Stack {
	return Stack{}
}

func (this *Stack) Push(x int) {
	this.elements = append(this.elements, x)
}

func (this *Stack) Pop() int {
	element := this.elements[this.Size()-1]
	this.elements = this.elements[:this.Size()-1]
	return element
}

func (this *Stack) Peek() int {
	return this.elements[this.Size()-1]
}

func (this *Stack) Empty() bool {
	if len(this.elements) != 0 {
		return false
	}
	return true
}

func (this *Stack) Size() int {
	return len(this.elements)
}
