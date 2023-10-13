package search

import "golang.org/x/exp/slices"

func BinarySearch(nums []int, searchData int) (int, bool) {
	var isFound bool
	var index int
	itemsCount := len(nums)
	slices.Sort(nums)
	for start, end := 0, itemsCount - 1 ; start <= end; {
		mid := (start + end) / 2

		if nums[mid] == searchData {
			index = mid
			isFound = true
			break
		}
		
		if searchData < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
		
	}

	return index, isFound
}