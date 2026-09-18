package stacks

/*
Min Stack
- When called getMin it should return a minimum value in the stack
- Idea is we store 2 values, value and min value until that point

*/

type Pair struct {
	Val int
	Min int
}

type MinStack struct {
	Elements []Pair
}

func (this *MinStack) Push(value int) {
	if len(this.Elements) == 0 {
		this.Elements = append(this.Elements, Pair{Val: value, Min: value})
	} else {
		this.Elements = append(this.Elements, Pair{Val: value, Min: min(value, this.Elements[len(this.Elements)-1].Min)})
	}
}

func (this *MinStack) Pop() {
	this.Elements = this.Elements[:len(this.Elements)-1]
}

func (this *MinStack) Top() int {
	return this.Elements[len(this.Elements)-1].Val
}

func (this *MinStack) GetMin() int {
	return this.Elements[len(this.Elements)-1].Min
}
