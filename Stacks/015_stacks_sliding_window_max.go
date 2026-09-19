package stacks

/* ( HARD to understand )
Sliding window maximum with window size as k
- let's consider small thing x1, x2, x3, x4 and window size is 3
- Let's bring an idea, is x1 is smaller than x2 is it really require to conside x1 in our result no we pop the element x1 and store x2
- Let's say if x1 is larger then we need to place x1 in queue and also x2 because x2 can be large for another sliding window as x1 cannot be
- That means while inserting we need to remove all lesser elements from the queue and insert x4 when shifting

- Now our front has the max value, but when we try to move forward by one bit then also we should remove front if the front idx <= i - k which means that it doesn't belong to the current window

*/

type Queue struct {
	Elements []int
}

func (q *Queue) PushBack(key int) {
	q.Elements = append(q.Elements, key)
}

func (q *Queue) PopFront() int {
	front := q.Elements[0]
	if len(q.Elements) == 0 {
		q.Elements = []int{}
	} else {
		q.Elements = q.Elements[1:]
	}
	return front
}

func (q *Queue) PopBack() int {
	back := q.Elements[len(q.Elements)-1]
	q.Elements = q.Elements[:len(q.Elements)-1]
	return back
}

func (q *Queue) Front() int {
	return q.Elements[0]
}

func (q *Queue) Back() int {
	return q.Elements[len(q.Elements)-1]
}

func (q *Queue) Empty() bool {
	return len(q.Elements) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	dQueue := Queue{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if !dQueue.Empty() && dQueue.Front() <= i-k {
			dQueue.PopFront()
		}

		for !dQueue.Empty() && nums[dQueue.Back()] <= nums[i] {
			dQueue.PopBack()
		}

		dQueue.PushBack(i)

		if i >= k-1 {
			res = append(res, nums[dQueue.Front()])
		}
	}

	return res
}
