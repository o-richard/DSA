package sort

func ShellSort(nums []int, isAsc bool) []int {
	itemsCount := len(nums)
	for interval := itemsCount / 2; interval > 0; interval /= 2 {
		for i := interval; i < itemsCount; i++ {
			temp := nums[i]
			j := i
			for ; isAsc && j >= interval && nums[j - interval] > temp; j -= interval {
				nums[j] = nums[j - interval]
			}
			for ; !isAsc && j >= interval && nums[j - interval] < temp; j -= interval {
				nums[j] = nums[j - interval]
			}
			nums[j] = temp
		}
	}

	return nums
}
