package stacks

/*
Minimum Span

- We have to figure out how many consecutive values are less than or equal to current price in before days
- 1, 2, 3, 4, 3, 6
day - 1:
[1 , 1]
day - 2:
[2, 2]
day - 3:
[3, 3]
day - 4:
[3, 3]
[3, 1]
day - 5:
[6, 5]

That means you are just clubbing the spans with <currPrice, total_spans_less_equal_price>
This works because single single value is being inserted
*/

type Stack016 struct {
	Elements [][2]int
}

func (s *Stack016) Push(key [2]int) {
	s.Elements = append(s.Elements, key)
}

func (s *Stack016) Pop() [2]int {
	front := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return front
}

func (s *Stack016) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack016) Top() [2]int {
	return s.Elements[len(s.Elements)-1]
}

type StockSpanner struct {
	PriceStack Stack016
}

func Constructor1() StockSpanner {
	return StockSpanner{
		PriceStack: Stack016{},
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for !this.PriceStack.Empty() && this.PriceStack.Top()[0] <= price {
		span += this.PriceStack.Pop()[1]
	}

	this.PriceStack.Push([2]int{price, span})
	return span

}
