package arrays

/*
Rearrange Positives and negatives but here count of pos != neg , start with pos but if remaining anything you can add at the end
*/

func rearrangeArray2(nums []int) []int {
	posArr := []int{}
	negArr := []int{}

	for _, val := range nums {
		if val > 0 {
			posArr = append(posArr, val)
		} else {
			negArr = append(negArr, val)
		}
	}

	pos := 0
	neg := 0
	insertPos := true
	res := []int{}

	for pos < len(posArr) && neg < len(negArr) {
		if insertPos {
			res = append(res, posArr[pos])
			pos++
		} else {
			res = append(res, negArr[neg])
			neg++
		}
		insertPos = !insertPos
	}

	for pos < len(posArr) {
		res = append(res, posArr[pos])
		pos++
	}

	for neg < len(negArr) {
		res = append(res, negArr[neg])
		neg++
	}

	return res

}
