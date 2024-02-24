func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	maxCandies := 0
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	result := make([]bool, n)
	for i, candy := range candies {
		if candy+extraCandies >= maxCandies {
			result[i] = true
		}
	}

	return result
}