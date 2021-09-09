package leetcode

func makeEqual(words []string) bool {
	dir := map[uint8]int{}

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			dir[words[i][j]] += 1
		}
	}

	for _, v := range dir {
		if v%len(words) != 0 {
			return false
		}
	}

	return true

}
