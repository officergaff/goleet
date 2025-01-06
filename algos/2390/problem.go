package algos

func removeStars(s string) string {
	var stack []byte
	for _, val := range s {
		if val == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(val))
		}
	}
	return string(stack)
}
