package trie

/*
Maximize Xor under some constraints ( value which are choosing aginst xor should be less than query[1])
- let's say If i choose want bit, i haven't got the result then i need to check not want bit which is bit
- if ctrNum > maxNum that means condition is invalid and i should return false
*/

func maxXorAdv(trie *BitTrie, idx int, ctrNum int, num int, maxNum int) (int, bool) {

	if ctrNum > maxNum {
		return -1, false
	}
	if idx < 0 {
		return 0, true
	}

	bit := (num >> idx) & 1
	want := bit ^ 1

	res := 0
	valid := false

	if trie.node[want] != nil {
		res, valid := maxXorAdv(trie.node[want], idx-1, ctrNum|(want<<idx), num, maxNum)
		if valid {
			return (1 << idx) | res, true
		}
	}

	if trie.node[bit] != nil {
		res, valid = maxXorAdv(trie.node[bit], idx-1, ctrNum|(bit<<idx), num, maxNum)
		if valid {
			return res, true
		}
	}

	return -1, false

}

func maximizeXor(nums []int, queries [][]int) []int {
	trie := BitTrie{
		node: [2]*BitTrie{},
	}
	for _, val := range nums {
		insert(&trie, 31, val)
	}

	result := []int{}

	for _, query := range queries {
		res, _ := maxXorAdv(&trie, 31, 0, query[0], query[1])
		result = append(result, res)
	}

	return result
}
