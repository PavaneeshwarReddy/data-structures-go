package queues

/*
Queue:
Why this because we need to maintain queue property that pushing happens from back and popping from front
Push : Append element at the end of list and pop front elements continously and move those elements behing curr
Pop: Current at idx 0 will be the top element and remove this
*/

type MyStack struct {
	Elements []int
}

func Constructor() MyStack {
	return MyStack{
		Elements: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.Elements = append(this.Elements, x)

	n := len(this.Elements) - 1

	for i := 0; i < n; i++ {
		front := this.Elements[0]
		this.Elements = this.Elements[1:]
		this.Elements = append(this.Elements, front)
	}
}

func (this *MyStack) Pop() int {
	top := this.Elements[0]
	this.Elements = this.Elements[1:]
	return top
}

func (this *MyStack) Top() int {
	return this.Elements[0]
}

func (this *MyStack) Empty() bool {
	return len(this.Elements) == 0
}
