package algos

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var greatest int
	for _, val := range candies {
		greatest = max(greatest, val)
	}
	var out []bool = make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= greatest {
			out[i] = true
		}
	}
	return out
}
