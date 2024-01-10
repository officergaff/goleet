package algos

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	min_len := min(len(word1), len(word2))
	var chars []string
	for i := 0; i < min_len; i++ {
		chars = append(chars, string(word1[i]), string(word2[i]))
	}
	if len(word1) > len(word2) {
		chars = append(chars, word1[min_len:])
	} else if len(word2) > len(word1) {
		chars = append(chars, word2[min_len:])
	}
	return strings.Join(chars, "")
}
