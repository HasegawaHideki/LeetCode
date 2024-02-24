func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		var current int = flowerbed[i]
		if current == 1 {
			continue
		}

		var prev, next int = 0, 0

		if i > 0 {
			prev = flowerbed[i-1]
		}

		if i < len(flowerbed)-1 {
			next = flowerbed[i+1]
		}

		if prev == 0 && next == 0 {
			count += 1
			flowerbed[i] = 1
		}
	}

	if count >= n {
		return true
	}

	return false
}