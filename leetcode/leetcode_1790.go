package leetcode

func areAlmostEqual(s1 string, s2 string) bool {
	var diff []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}
	if len(diff) == 0 {
		return true
	} else if len(diff) == 2 {
		return s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]]
	} else {
		return false
	}

}
