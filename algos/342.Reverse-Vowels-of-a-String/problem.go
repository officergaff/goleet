package algos

import (
	"strings"
)

func ReverseVowels(s string) string {
	arr := make([]string, len(s))
	l, r := 0, len(s)-1
	for l <= r {
		l_val, r_val := string(s[l]), string(s[r])
		if isVowel(l_val) && isVowel(r_val) {
			arr[l] = string(s[r])
			arr[r] = string(s[l])
			l += 1
			r -= 1
			continue
		}
		if !isVowel(l_val) {
			arr[l] = l_val
			l += 1
		}
		if !isVowel(r_val) {
			arr[r] = r_val
			r -= 1
		}
	}
	return strings.Join(arr, "")
}

func isVowel(s string) bool {
	switch s {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		return true
	}
	return false
}
