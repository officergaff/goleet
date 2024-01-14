package algos

import "strings"

func ReverseWords(s string) string {
	words := strings.Fields(s)
	arr := make([]string, len(words))
	l, r := 0, len(words)-1
	for l <= r {
		arr[l], arr[r] = words[r], words[l]
		l += 1
		r -= 1
	}
	return strings.Join(arr, " ")
}
