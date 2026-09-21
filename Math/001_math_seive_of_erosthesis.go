package math

/*
1. We take the number and mark all multiples to false
2. We only process if it's already a prime or else skip them
*/

func countPrimes(n int) int {
	isPrimes := make([]bool, n)
	for i := range n {
		if i > 1 {
			isPrimes[i] = true
		}
	}

	for i := 2; i*i < n; i++ {
		if isPrimes[i] {
			for j := i * i; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}

	res := 0
	for _, val := range isPrimes {
		if val {
			res += 1
		}
	}

	return res

}
