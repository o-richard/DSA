package search

func LinearSearch(nums []int, searchData int) bool {
	var isFound bool
	for _, num := range nums {
		if searchData == num {
			isFound = true
			break
		}
	}
	return isFound
}