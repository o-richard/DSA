package sort

func merge(nums []int, start, mid, end int, isAsc bool) {
	size1 := (mid - start) + 1
	size2 := end - mid

	arr1 := make([]int, size1)
	arr2 := make([]int, size2)

	for i := 0; i < size1; i++ {
		arr1[i] = nums[start + i]
	}
	for i := 0; i < size2; i++ {
		arr2[i] = nums[mid + i + 1]
	}

	var start1, start2 int
	originalStart := start

	for start1 < size1 && start2 < size2 {
		if isAsc {
			if arr1[start1] <= arr2[start2] {
				nums[originalStart] = arr1[start1]
				start1++
			} else {
				nums[originalStart] = arr2[start2]
				start2++
			}
		} else {
			if arr1[start1] > arr2[start2] {
				nums[originalStart] = arr1[start1]
				start1++
			} else {
				nums[originalStart] = arr2[start2]
				start2++
			}
		}
		originalStart++
	}

	for i := start1; i < size1; i++ {
		nums[originalStart] = arr1[i]
		originalStart++
	}
	for i := start2; i < size2; i++ {
		nums[originalStart] = arr2[i]
		originalStart++
	}

}

func mergeSortAlg(nums []int, start, end int, isAsc bool) {
	if start < end {
		mid := (start + end) / 2
		mergeSortAlg(nums, start, mid, isAsc)
		mergeSortAlg(nums, mid + 1, end, isAsc)
		merge(nums, start, mid, end, isAsc)
	}
}

func MergeSort(nums []int, isAsc bool) []int {
	mergeSortAlg(nums, 0, len(nums) - 1, isAsc)
	return nums
}