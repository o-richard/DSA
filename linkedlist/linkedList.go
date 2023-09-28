package linkedlist

type linkedList interface {
	tranverse(goLeftToRight bool) []int
	search(searchData int) (int, bool)
	insertion(inputData int, insertionIndex int) error
	retrieval(retrievalIndex int, isDelete bool) (int, error)
}

func Tranverse(ll linkedList, goLeftToRight bool) []int {
	return ll.tranverse(goLeftToRight)
}

func Search(ll linkedList, searchData int) (int, bool) {
	return ll.search(searchData)
}

func Insertion(ll linkedList, inputData int, insertionIndex int) error {
	return ll.insertion(inputData, insertionIndex)
}

func Retrieval(ll linkedList, retrievalIndex int, isDelete bool) (int, error) {
	return ll.retrieval(retrievalIndex, isDelete)
}