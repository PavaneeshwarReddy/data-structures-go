package heaps

/*
Finding median

- Consider this we need one value from the left and one value right half if length is EVEN
- That means highest value from the left half and lowest value from right half
- That means left will be the max heap and right will be the min heap
- We will insert value in left if empty or top is greater than num
- But later we need rebalance as we can accept only +1 size difference not more than that
*/

type MedianFinder struct {
	rightHeap MinHeap
	leftHeap  MaxHeap
}

func Constructor2() MedianFinder {
	return MedianFinder{rightHeap: MinHeap{}, leftHeap: MaxHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	if this.leftHeap.Length() == 0 || num <= this.leftHeap.Peek() {
		this.leftHeap.Insert(num)
	} else {
		this.rightHeap.Insert(num)
	}

	if this.leftHeap.Length() > this.rightHeap.Length()+1 {
		this.rightHeap.Insert(this.leftHeap.Pop())
	}

	if this.rightHeap.Length() > this.leftHeap.Length() {
		this.leftHeap.Insert(this.rightHeap.Pop())
	}

}

func (this *MedianFinder) FindMedian() float64 {

	if this.leftHeap.Length() != this.rightHeap.Length() {
		return float64(this.leftHeap.Peek())
	}

	return (float64(this.leftHeap.Peek()) + float64(this.rightHeap.Peek())) / 2

}
