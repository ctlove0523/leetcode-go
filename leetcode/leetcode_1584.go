package leetcode

func maximumPopulation(logs [][]int) int {
	years := make([]int, 101)

	for i := 0; i < len(logs); i++ {
		log := logs[i]
		years[log[0]-1950] += 1
		years[log[1]-1950] -= 1
	}

	maxPeople := years[0]
	maxYear := 0
	for i := 1; i < len(years); i++ {
		years[i] = years[i] + years[i-1]
		if years[i] > maxPeople {
			maxPeople = years[i]
			maxYear = i
		}
	}

	return maxYear + 1950

}
