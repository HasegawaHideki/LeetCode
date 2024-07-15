func removeDuplicates(nums []int) int {
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[index] == nums[i] {
			nums[i] = 999
		} else {
			index = i
		}
	}

	index = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 999 {
			tmp := nums[index]
			nums[index] = nums[i]
			nums[i] = tmp
			index++
		}
	}

	return index
}