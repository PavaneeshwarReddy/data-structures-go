package arrays

/*
Union of two sorted arrays
- There can be other approaches
- Set, Map
- We will solve using 2 pointers
*/

func FindUnion(num1 []int, num2 []int) []int {
	union := []int{}

	i, j := 0, 0

	for i < len(num1) && j < len(num2) {
		if num1[i] < num2[j] {
			if len(union) == 0 || union[len(union)-1] != num1[i] {
				union = append(union, num1[i])
			}
			i++
		} else if num2[j] < num1[i] {
			if len(union) == 0 || union[len(union)-1] != num2[j] {
				union = append(union, num2[j])
			}
			j++
		} else {
			if len(union) == 0 || union[len(union)-1] != num2[j] {
				union = append(union, num2[j])
			}
			j++
			i++
		}
	}

	for i < len(num1) {
		if len(union) == 0 || union[len(union)-1] != num1[i] {
			union = append(union, num1[i])
		}
		i++
	}

	for j < len(num2) {
		if len(union) == 0 || union[len(union)-1] != num2[j] {
			union = append(union, num2[j])
		}
		j++
	}

	return union
}
