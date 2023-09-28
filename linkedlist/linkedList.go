package linkedlist

type LinkedList interface {
	tranverse(goLeftToRight bool) []int
	search(searchData int) (int, bool)
	insertion(inputData int, insertionIndex int) error
	retrieval(retrievalIndex int, isDelete bool) (int, error)
}

func Tranverse(ll LinkedList, goLeftToRight bool) []int {
	return ll.tranverse(goLeftToRight)
}

func Search(ll LinkedList, searchData int) (int, bool) {
	return ll.search(searchData)
}

func Insertion(ll LinkedList, inputData int, insertionIndex int) error {
	return ll.insertion(inputData, insertionIndex)
}

func Retrieval(ll LinkedList, retrievalIndex int, isDelete bool) (int, error) {
	return ll.retrieval(retrievalIndex, isDelete)
}