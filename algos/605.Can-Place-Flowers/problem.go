package algos

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 1 {
			continue
		}
		var prev int
		if i != 0 {
			prev = flowerbed[i-1]
		}
		var next int
		if i != len(flowerbed)-1 {
			next = flowerbed[i+1]
		}
		if prev == 0 && next == 0 {
			flowerbed[i] = 1
			n -= 1
		}
	}
	return n == 0
}
