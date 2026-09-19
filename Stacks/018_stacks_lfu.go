package stacks

/*
LFU cache

- Consider MinFreq is 0 at the point intially

- Now let's say you are inserting a new element which doesn't exisits and its the start, doesnt exisits in KeyNode
	- Check whether DLL exisits with the required Freq of the node which is 1
	- If not create a new DLL and add this particular node to that DLL at the beginning I mean near the head
	- Now increase the MinFreq to 1 as it is the first one

- Now let's say you are going to put the same key with different value
	- Pick the node from keyNode
	- Consider old frequency, remove the node from the DLL of it's previous frequency
	- Check whether old DLL is out of nodes or not, if out of nodes then delete that DLL also, if the freq is same as old node then increment MinFreq also
	- Add it to new frequency DLL


- Node let's say you are inserting a new key with the capacity exceeded
	- Check the min frequency, remove last element from there
	- if that list is empty simply delete the DLL node
	- Add this into minFreq 1 DLL

*/

type DLLNode018 struct {
	Key  int
	Val  int
	Freq int
	Next *DLLNode018
	Prev *DLLNode018
}

type DLL018 struct {
	Head *DLLNode018
	Tail *DLLNode018
}

type LFUCache struct {
	KeyNode  map[int]*DLLNode018
	FreqList map[int]*DLL018
	MinFreq  int
	Capacity int
}

func Constructor018(capacity int) LFUCache {
	return LFUCache{
		KeyNode:  make(map[int]*DLLNode018),
		FreqList: make(map[int]*DLL018),
		MinFreq:  0,
		Capacity: capacity,
	}
}

func (this *LFUCache) createFreqList(freq int) {
	head := &DLLNode018{
		Key: -1,
		Val: -1,
	}

	tail := &DLLNode018{
		Key: -1,
		Val: -1,
	}

	head.Next = tail
	tail.Prev = head

	this.FreqList[freq] = &DLL018{
		Head: head,
		Tail: tail,
	}
}

func (this *LFUCache) addToFront(freq int, node *DLLNode018) {
	dll, ok := this.FreqList[freq]

	if !ok {
		this.createFreqList(freq)
		dll = this.FreqList[freq]
	}

	first := dll.Head.Next

	dll.Head.Next = node
	node.Prev = dll.Head

	node.Next = first
	first.Prev = node
}

func (this *LFUCache) removeNode(node *DLLNode018) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	node.Prev = nil
	node.Next = nil
}

func (this *LFUCache) removeLast(freq int) *DLLNode018 {
	dll := this.FreqList[freq]

	node := dll.Tail.Prev

	if node == dll.Head {
		return nil
	}

	this.removeNode(node)

	return node
}

func (this *LFUCache) increaseFreq(node *DLLNode018) {
	oldFreq := node.Freq

	this.removeNode(node)

	dll := this.FreqList[oldFreq]

	if dll.Head.Next == dll.Tail {
		delete(this.FreqList, oldFreq)

		if this.MinFreq == oldFreq {
			this.MinFreq++
		}
	}

	node.Freq++

	this.addToFront(node.Freq, node)
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.KeyNode[key]

	if !ok {
		return -1
	}

	this.increaseFreq(node)

	return node.Val
}

func (this *LFUCache) Put(key int, value int) {
	if this.Capacity == 0 {
		return
	}

	if node, ok := this.KeyNode[key]; ok {
		node.Val = value

		this.increaseFreq(node)

		return
	}

	if len(this.KeyNode) >= this.Capacity {
		node := this.removeLast(this.MinFreq)

		delete(this.KeyNode, node.Key)

		dll := this.FreqList[this.MinFreq]

		if dll.Head.Next == dll.Tail {
			delete(this.FreqList, this.MinFreq)
		}
	}

	node := &DLLNode018{
		Key:  key,
		Val:  value,
		Freq: 1,
	}

	this.KeyNode[key] = node

	this.addToFront(1, node)

	this.MinFreq = 1
}
