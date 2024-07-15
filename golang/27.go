func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			tmp := nums[index]
			nums[index] = nums[i]
			nums[i] = tmp
			index++
		}
	}

	return index
}