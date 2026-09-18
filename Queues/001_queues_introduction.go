package queues

import "errors"

/*
Queue

PushBack -> Pushed back of the list
PopFront -> Pop element from the front of list
Peek -> Gives the top element
Rear -> Gives the last element

ErrQueueEmpty -> Queue is empty
*/

var ErrQueueEmpty = errors.New("Queue is empty")

type Queue struct {
	Elements []int
}

func (q *Queue) PushBack(ele int) {
	q.Elements = append(q.Elements, ele)
}

func (q *Queue) PopFront() (int, error) {
	if len(q.Elements) == 0 {
		return -1, ErrQueueEmpty
	}
	frontEle := q.Elements[0]
	if len(q.Elements) == 1 {
		q.Elements = []int{}
	} else {
		q.Elements = q.Elements[1:]
	}
	return frontEle, nil
}

func (q *Queue) Peek() (int, error) {
	if len(q.Elements) == 0 {
		return -1, ErrQueueEmpty
	}
	return q.Elements[0], nil
}

func (q *Queue) Rear() (int, error) {
	if len(q.Elements) == 0 {
		return -1, ErrQueueEmpty
	}
	return q.Elements[len(q.Elements)-1], nil
}
