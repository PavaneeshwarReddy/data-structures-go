package trie

/*
Maximum Xor of 2 numbers in an array
- In normal trie we have 26 characters but in bit trie we only have 2
- We need to follow greedy algo, if our bit is 1 we have 0 then choose 0, else unset the bit
- Return max for every number
- Complexity is minimal as O ( 32 * n )
*/

type BitTrie struct {
	node [2]*BitTrie
}

func insert(trie *BitTrie, bit int, num int) {
	if bit < 0 {
		return
	}

	decidingBit := int((num >> bit) & 1)
	if trie.node[decidingBit] == nil {
		newTrie := BitTrie{
			node: [2]*BitTrie{},
		}
		trie.node[decidingBit] = &newTrie
	}

	insert(trie.node[decidingBit], bit-1, num)
}

func maxXor(trie *BitTrie, bit int, num int) int {
	if bit < 0 {
		return 0
	}

	b := (num >> bit) & 1
	want := b ^ 1

	if trie.node[want] != nil {
		return (1 << bit) | maxXor(trie.node[want], bit-1, num)
	}

	return maxXor(trie.node[b], bit-1, num)
}

func findMaximumXOR(nums []int) int {
	trie := &BitTrie{}

	for _, num := range nums {
		insert(trie, 31, num)
	}

	result := 0

	for _, num := range nums {
		result = max(result, maxXor(trie, 31, num))
	}

	return result
}
