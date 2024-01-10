package algos

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	var largest, smallest string
	if len(str1) >= len(str2) {
		largest = str1
		smallest = str2
	} else {
		largest = str2
		smallest = str1
	}
	for i := len(smallest); i >= 0; i-- {
		ss := smallest[:i]
		if ss != "" && len(largest)%len(ss) == 0 && len(smallest)%len(ss) == 0 {
			if strings.Repeat(ss, len(largest)/len(ss)) == largest && strings.Repeat(ss, len(smallest)/len(ss)) == smallest {
				return ss
			}
		}
	}
	return ""
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func gcdOfStringsOptimized(s1 string, s2 string) string {
	if s1+s2 != s2+s1 {
		return ""
	}
	x := gcd(len(s1), len(s2))
	return s1[:x]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
