package heaps

/*
Hands of Straights
- hand = [1,2,3,6,2,3,4,7,8], groupSize = 3, can we divide into 3 groups such that those groups contains consecutive card number
- We can maintain a minHeap with card number
- Pop the card and try to pop the next min card if diff is not more than 1 then its consecutive
- Decrement count by 1 and process. Add to the queue because including this won't give use results
- At last after processing each group add processed elements back
*/

type Card struct {
	Num   int
	Count int
}

type CardMinHeap struct {
	nodes []Card
}

func (cmh *CardMinHeap) Insert(node Card) {
	cmh.nodes = append(cmh.nodes, node)

	idx := len(cmh.nodes) - 1

	for idx > 0 {
		parentIdx := (idx - 1) / 2

		if parentIdx >= 0 && cmh.nodes[parentIdx].Num > cmh.nodes[idx].Num {
			cmh.nodes[parentIdx], cmh.nodes[idx] = cmh.nodes[idx], cmh.nodes[parentIdx]
		}

		idx = parentIdx
	}
}

func (cmh *CardMinHeap) Pop() Card {
	peekNode := cmh.nodes[0]
	cmh.nodes[0] = cmh.nodes[len(cmh.nodes)-1]
	cmh.nodes = cmh.nodes[:len(cmh.nodes)-1]

	idx := 0

	for true {
		leftChild := 2*idx + 1
		rightChild := 2*idx + 2
		smallest := idx

		if leftChild < len(cmh.nodes) && cmh.nodes[leftChild].Num < cmh.nodes[smallest].Num {
			smallest = leftChild
		}

		if rightChild < len(cmh.nodes) && cmh.nodes[rightChild].Num < cmh.nodes[smallest].Num {
			smallest = rightChild
		}

		if smallest != idx {
			cmh.nodes[smallest], cmh.nodes[idx] = cmh.nodes[idx], cmh.nodes[smallest]
			idx = smallest
		} else {
			break
		}
	}
	return peekNode
}

func (cmh *CardMinHeap) Length() int {
	return len(cmh.nodes)
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	minHeap := CardMinHeap{}
	freqMap := make(map[int]int)
	for _, val := range hand {
		_, ok := freqMap[val]
		if !ok {
			freqMap[val] = 1
		} else {
			freqMap[val] += 1
		}
	}

	for k, v := range freqMap {
		minHeap.Insert(Card{Num: k, Count: v})
	}

	for minHeap.Length() > 0 {
		currGroup := groupSize - 1
		card := minHeap.Pop()
		queue := []Card{}
		if card.Count != 1 {
			card.Count -= 1
			queue = append(queue, card)
		}

		for currGroup > 0 {
			if minHeap.Length() == 0 {
				return false
			}
			currCard := minHeap.Pop()
			if currCard.Num-1 != card.Num {
				return false
			}
			if currCard.Count != 1 {
				currCard.Count -= 1
				queue = append(queue, currCard)
			}
			currGroup -= 1
			card = currCard
		}
		for _, val := range queue {
			minHeap.Insert(val)
		}
	}

	return true

}
