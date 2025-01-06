package algos

func shiftingLetters(s string, shifts []int) string {
	runes := []rune{}
	for i := len(shifts) - 2; i >= 0; i-- {
		// This avoids int overflow
		shifts[i] += shifts[i+1] % 26
	}
	for idx, c := range s {
		cc := shift(c, shifts[idx])
		runes = append(runes, cc)
	}
	return string(runes)
}

func shift(s rune, shift int) rune {
	return (s-'a'+rune(shift))%26 + 'a'
}
