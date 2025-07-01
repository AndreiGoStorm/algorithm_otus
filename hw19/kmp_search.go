package hw19

type KMPSearch struct {
	pattern string
	pi      []int
}

func NewKMPSearch(pattern string) *KMPSearch {
	kmp := &KMPSearch{pattern: pattern}
	kmp.createPi(pattern)
	return kmp
}

func (kmp *KMPSearch) createPi(pattern string) {
	kmp.pi = make([]int, len(pattern)+1)
	for i := 1; i < len(pattern); i++ {
		length := kmp.pi[i]
		for length > 0 && pattern[length] != pattern[i] {
			length = kmp.pi[length]
		}
		if pattern[length] == pattern[i] {
			length++
		}
		kmp.pi[i+1] = length
	}
}

func (kmp *KMPSearch) Search(text string) []int {
	kmp.createPi(kmp.pattern + "#" + text)
	res := make([]int, 0)
	for i := 0; i < len(kmp.pi); i++ {
		if kmp.pi[i] == len(kmp.pattern) {
			res = append(res, i-len(kmp.pattern)*2-1)
		}
	}
	return res
}
