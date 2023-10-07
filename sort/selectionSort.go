package sort

func SelectionSort(nums []int, isAsc bool) []int {
	for i := range(nums) {
		for j := range(nums[i:]) {
			correctIndex := j + i
			if isAsc {
				if nums[correctIndex] < nums[i] {
					nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
				}
			} else {
				if nums[correctIndex] > nums[i] {
					nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
				}
			}
		}
	}
	return nums
}