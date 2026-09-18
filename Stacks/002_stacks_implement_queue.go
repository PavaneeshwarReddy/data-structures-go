package stacks

/*
Queue using Stack

Push: When inserted we need to pop all and store in new stack append this and push all those while popping out from stack
Peek: Return the top element because as we rotated stack
*/

type MyQueue struct {
	Elements []int
}

func Constructor() MyQueue {
	return MyQueue{
		Elements: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	newStack := []int{}
	n := len(this.Elements)
	for _ = range n {
		newStack = append(newStack, this.Elements[len(this.Elements)-1])
		this.Elements = this.Elements[:len(this.Elements)-1]
	}
	this.Elements = append(this.Elements, x)
	for _ = range n {
		this.Elements = append(this.Elements, newStack[len(newStack)-1])
		newStack = newStack[:len(newStack)-1]
	}
}

func (this *MyQueue) Pop() int {
	top := this.Elements[len(this.Elements)-1]
	this.Elements = this.Elements[:len(this.Elements)-1]
	return top
}

func (this *MyQueue) Peek() int {
	return this.Elements[len(this.Elements)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.Elements) == 0
}
