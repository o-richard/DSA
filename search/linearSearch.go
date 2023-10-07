package search

func LinearSearch(nums []int, searchData int) (int, bool) {
	var isFound bool
	var index int
	for i, num := range nums {
		if searchData == num {
			index = i
			isFound = true
			break
		}
	}
	return index, isFound
}