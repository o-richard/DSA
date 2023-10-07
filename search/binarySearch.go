package search

func maxHeapify(nums []int, itemsCount, largestIndex int) {
	largest := largestIndex
	left := 2 * largestIndex + 1
	right := 2 * largestIndex + 2

	if left < itemsCount && nums[largestIndex] < nums[left] {
		largest = left
	}

	if right < itemsCount && nums[largest] < nums[right] {
		largest = right
	}

	if largest != largestIndex {
		nums[largestIndex], nums[largest] = nums[largest], nums[largestIndex]
		maxHeapify(nums, itemsCount, largest)
	}
}

func heapSort(nums []int) ([]int, int) {
	itemsCount := len(nums)
	for i := (itemsCount / 2); i > -1; i-- {
		maxHeapify(nums, itemsCount, i)
	}
	for i := itemsCount - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		maxHeapify(nums, i, 0)
	}
	return nums, itemsCount
} 

func BinarySearch(nums []int, searchData int) (int, bool) {
	var isFound bool
	var index int
	nums, itemsCount := heapSort(nums)
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