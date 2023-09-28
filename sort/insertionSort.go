package sort

func InsertionSort(nums []int, isAsc bool) []int {
	itemsCount := len(nums)
	for i := 1; i < itemsCount; i++ {
		key := nums[i]
		j := i - 1

		for isAsc && j >= 0 && key < nums[j] {
			nums[j + 1] = nums[j]
			j--
		}
		
		for !isAsc && j >= 0 && key > nums[j] {
			nums[j + 1] = nums[j]
			j--
		}

		nums[j + 1] = key
	}
	return nums
}