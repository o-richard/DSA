package sort

func partition(nums []int, start, end int, isAsc bool) int {
	pivot := nums[end]
	root := start - 1

	for i := start; i < end; i++ {
		if isAsc && nums[i] <= pivot {
			root++
			nums[root], nums[i] = nums[i], nums[root]
		}
		if !isAsc && nums[i] >= pivot {
			root++
			nums[root], nums[i] = nums[i], nums[root]
		}
	}
	
	nums[root + 1], nums[end] = nums[end], nums[root + 1]
	return root + 1
}

func quickSortAlg(nums []int, start, end int, isAsc bool) {
	if start < end {
		pivot := partition(nums, start, end, isAsc)
		quickSortAlg(nums, start, pivot - 1, isAsc)
		quickSortAlg(nums, pivot + 1, end, isAsc)
	}
}

func QuickSort(nums []int, isAsc bool) []int {
	quickSortAlg(nums, 0, len(nums) - 1, isAsc)
	return nums
}