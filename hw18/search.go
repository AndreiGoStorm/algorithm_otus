package hw18

func fullScanSearch(text, mask string) (int, int) {
	runeText, runeMask := []rune(text), []rune(mask)
	var t, cmp int
	for t <= len(runeText)-len(runeMask) {
		m := 0
		for m < len(runeMask) && runeText[t+m] == runeMask[m] {
			cmp++
			m++
		}
		if m == len(runeMask) {
			return t + 1, cmp
		}
		t++
	}
	return -1, cmp
}

func backScanSearch(text, mask string) (int, int) {
	runeText, runeMask := []rune(text), []rune(mask)
	var t, cmp int
	for t <= len(runeText)-len(runeMask) {
		m := len(runeMask) - 1
		for m >= 0 && runeText[t+m] == runeMask[m] {
			cmp++
			m--
		}
		if m < 0 {
			return t + 1, cmp
		}
		t++
	}
	return -1, cmp
}

func boyerMooreHorspoolShift1(text, mask string) (int, int) {
	runeText, runeMask := []rune(text), []rune(mask)
	shift := prepareBMH1(runeMask)
	var t, cmp int
	for t <= len(runeText)-len(runeMask) {
		m := len(runeMask) - 1
		for m >= 0 && runeText[t+m] == runeMask[m] {
			cmp++
			m--
		}
		if m < 0 {
			return t + 1, cmp
		}
		if _, ok := shift[runeText[t+m]]; ok {
			t++
			continue
		}
		t += len(runeMask) - 1
	}
	return -1, cmp
}

func prepareBMH1(mask []rune) map[rune]struct{} {
	shift := make(map[rune]struct{}, len(mask))
	for _, r := range mask {
		shift[r] = struct{}{}
	}

	return shift
}

func boyerMooreHorspool(text, mask string) (int, int) {
	runeText, runeMask := []rune(text), []rune(mask)
	shift := prepareBMH(runeMask)
	var t, cmp int
	for t <= len(runeText)-len(runeMask) {
		m := len(runeMask) - 1
		for m >= 0 && runeText[t+m] == runeMask[m] {
			cmp++
			m--
		}
		if m < 0 {
			return t + 1, cmp
		}
		if value, ok := shift[runeText[t+m]]; ok {
			t += value
			continue
		}
		t += len(runeMask) - 1
	}
	return -1, cmp
}

func prepareBMH(mask []rune) map[rune]int {
	shift := make(map[rune]int, len(mask))
	for i := 0; i < len(mask)-1; i++ {
		shift[mask[i]] = len(mask) - i - 1
	}

	return shift
}

func boyerMoore(text, mask string) (int, int) {
	runeText, runeMask := []rune(text), []rune(mask)
	suffixes := createSuffixTable(runeMask)
	var t, cmp int
	for t <= len(runeText)-len(runeMask) {
		m := len(runeMask) - 1
		for m >= 0 && runeText[t+m] == runeMask[m] {
			cmp++
			m--
		}
		if m < 0 {
			return t + 1, cmp
		}
		t += suffixes[len(runeMask)-1-m]
	}
	return -1, cmp
}

func createSuffixTable(mask []rune) []int {
	suffixes := []int{1}
	for i := 1; i < len(mask); i++ {
		for k := 1; k <= len(mask); k++ {
			cnt := 0
			for j := 0; j < i; j++ {
				if j+k+1 > len(mask) {
					break
				}
				if mask[len(mask)-1-j] != mask[len(mask)-1-j-k] {
					break
				}
				cnt++
			}

			if cnt < i {
				if k+cnt == len(mask) {
					suffixes = append(suffixes, k)
					break
				}
			} else {
				suffixes = append(suffixes, k)
				break
			}
		}
	}

	return suffixes
}
