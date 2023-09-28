package search

import "slices"

func BinarySearch(nums []int, searchData int) bool {
	var isFound bool
	itemsCount := len(nums)
	slices.Sort(nums)
	for start, end := 0, itemsCount - 1 ; start <= end; {
		mid := (start + end) / 2

		if nums[mid] == searchData {
			isFound = true
			break
		}
		
		if searchData < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
		
	}

	return isFound
}