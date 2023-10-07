package sort

func BubbleSort(nums []int, isAsc bool) []int {
	itemsCount := len(nums)
	for i := 0; i < itemsCount; i++ {
		var swapped bool
		for j := 0; j < itemsCount - i - 1; j++ {
			if isAsc {
				if nums[j] > nums[j + 1] {
					nums[j], nums[j + 1] = nums[j + 1], nums[j]
					swapped = true
				}
			} else {
				if nums[j] < nums[j + 1] {
					nums[j], nums[j + 1] = nums[j + 1], nums[j]
					swapped = true
				}
			}
		}
		if !swapped {
			break
		}
	}
	return nums
}