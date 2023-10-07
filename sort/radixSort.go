package sort

import "errors"

func countingSort(nums []int, placeVal, itemCount int) {
	var maxElem int
	for i := 0; i < itemCount; i++ {
		val := (nums[i] / placeVal) % 10
		if val > maxElem {
			maxElem = val
		}
	}

	countArr := make([]int, maxElem + 1)
	for i := range nums {
		index := (nums[i] / placeVal) % 10
		countArr[index]++
	}

	var prevCount int
	for i := range countArr {
		countArr[i] += prevCount
		prevCount = countArr[i] 
	}

	outputArr := make([]int, itemCount)
	for i := itemCount - 1; i >= 0; i-- {
		index := (nums[i] / placeVal) % 10
		countArr[index]--
		outputArr[countArr[index]] = nums[i]
	}
	
	copy(nums, outputArr)
}

func RadixSort(nums []int) ([]int, error) {
	itemCount := len(nums) 
	
	var maxElem int
	for i := 0; i < itemCount; i++ {
		if nums[i] < 0 {
			return nums, errors.New("negative values are not supported")
		}
		if nums[i] > maxElem {
			maxElem = nums[i]
		}
	}

	for placeVal := 1; maxElem / placeVal > 0; placeVal *= 10 {
		countingSort(nums, placeVal, itemCount)
	}

	return nums, nil
}