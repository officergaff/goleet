package algos

/*
store the count of character inside of a map (num, occurrences)
iterate over the map, storing the (occurences, {})
*/
func UniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	for _, val := range arr {
		count[val] += 1
	}
	oMap := make(map[int]struct{})
	for _, val := range count {
		if _, ok := oMap[val]; ok {
			return false
		}
		oMap[val] = struct{}{}
	}
	return true
}
