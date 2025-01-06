package algos

/*
 */
func findDifference(nums1 []int, nums2 []int) [][]int {
	var n1, n2 []int
	m1, m2 := make(map[int]struct{}), make(map[int]struct{})
	for _, val := range nums1 {
		m1[val] = struct{}{}
	}
	for _, val := range nums2 {
		m2[val] = struct{}{}
	}
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			n1 = append(n1, k)
		}
	}
	for k := range m2 {
		if _, ok := m1[k]; !ok {
			n2 = append(n2, k)
		}
	}
	return [][]int{n1, n2}
}
