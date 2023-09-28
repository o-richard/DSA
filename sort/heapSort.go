package sort

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

func minHeapify(nums []int, itemsCount, smallestIndex int) {
	smallest := smallestIndex
	left := 2 * smallestIndex + 1
	right := 2 * smallestIndex + 2

	if left < itemsCount && nums[smallestIndex] > nums[left] {
		smallest = left
	}

	if right < itemsCount && nums[smallest] > nums[right] {
		smallest = right
	}

	if smallest != smallestIndex {
		nums[smallestIndex], nums[smallest] = nums[smallest], nums[smallestIndex]
		minHeapify(nums, itemsCount, smallest)
	}
	
}

func HeapSort(nums []int, isAsc bool) []int {
	itemsCount := len(nums)
	for i := (itemsCount / 2); i > -1; i-- {
		if isAsc {
			maxHeapify(nums, itemsCount, i)
		} else {
			minHeapify(nums, itemsCount, i)
		}
	}
	for i := itemsCount - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		if isAsc {
			maxHeapify(nums, i, 0)
		} else {
			minHeapify(nums, i, 0)
		}
	}
	return nums
} 