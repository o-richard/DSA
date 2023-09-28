package sort

import "errors"

func CountingSort(nums []int) ([]int, error) {
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
	
	countArr := make([]int, maxElem + 1)
	for i := 0; i < itemCount; i++ {
		countArr[nums[i]]++
	}
	
	var prevCount int
	for i := range countArr {
		countArr[i] += prevCount
		prevCount = countArr[i] 
	}
	
	outputArr := make([]int, itemCount)
	for _, num := range nums {
		countArr[num]--
		outputArr[countArr[num]] = num
	}

	return outputArr, nil
}