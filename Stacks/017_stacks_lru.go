package stacks

/*
LRU Cache
- Get -> if the value is present return the value else -1 and also the key should be moved forward as it used frequently
- Put -> if the key exisits overwrite and if its new and exceeds the capacity remove the LRU key from cache

- Elements near the head is the Recently used and near the tail it's not recently used
- When we are removing out of capacity then we need to remove near tail and add near head
*/
type DLLNode struct {
	Key  int
	Val  int
	Next *DLLNode
	Prev *DLLNode
}

type DLL struct {
	Head *DLLNode
	Tail *DLLNode
}

type LRUCache struct {
	KeyNode  map[int]*DLLNode
	DList    DLL
	Capacity int
}

func Constructor017(capacity int) LRUCache {
	head := DLLNode{Val: -1}
	tail := DLLNode{Val: -1}
	head.Next = &tail
	tail.Prev = &head

	return LRUCache{
		KeyNode: make(map[int]*DLLNode),
		DList: DLL{
			Head: &head,
			Tail: &tail,
		},
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.KeyNode[key]; ok {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		headNext := this.DList.Head.Next

		this.DList.Head.Next = node
		node.Prev = this.DList.Head
		node.Next = headNext
		headNext.Prev = node
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.KeyNode[key]; ok {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		headNext := this.DList.Head.Next
		node.Val = value
		this.DList.Head.Next = node
		node.Prev = this.DList.Head
		node.Next = headNext
		headNext.Prev = node
	} else {
		if len(this.KeyNode) >= this.Capacity {
			// find the least recently used and delete it
			deleteNode := this.DList.Tail.Prev
			deleteNode.Prev.Next = deleteNode.Next
			deleteNode.Next.Prev = deleteNode.Prev
			delete(this.KeyNode, deleteNode.Key)
		}

		// insert new node near the head
		newNode := DLLNode{
			Key: key,
			Val: value,
		}
		headNext := this.DList.Head.Next
		this.DList.Head.Next = &newNode
		newNode.Prev = this.DList.Head
		newNode.Next = headNext
		headNext.Prev = &newNode

		this.KeyNode[key] = &newNode
	}
}
