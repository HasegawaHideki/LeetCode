func increasingTriplet(nums []int) bool {
	num1 := math.MaxInt
	num2 := math.MaxInt

	for i := 0; i < len(nums); i++ {
		if nums[i] <= num1 {
			num1 = nums[i]
		} else if nums[i] <= num2 {
			num2 = nums[i]
		} else {
			return true
		}
	}

	return false
}