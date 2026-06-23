package strings

/*
Isomorphic Strings
	- Take two maps which maps to character index for both strings
	- When you fetch and their index is not equal then they are not isomorphic
*/

func Isomorphic(s string, t string) bool {
	mpS := make(map[string]int)
	mpT := make(map[string]int)

	for i := 0; i < len(s); i++ {
		valS, okS := mpS[string(s[i])]
		if !okS {
			mpS[string(s[i])] = i
			valS = i
		}

		valT, okT := mpT[string(t[i])]
		if !okT {
			mpT[string(t[i])] = i
			valT = i
		}

		if valS != valT {
			return false
		}
	}

	return true
}
