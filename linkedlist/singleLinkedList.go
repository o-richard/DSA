package linkedlist

import (
	"fmt"
	"errors"
)

// sllNode refers to a singleLinkedListNode
// SLL refers to Single Linked List

type sllNode struct {
	data int
	next *sllNode
}

type sLL struct {
	itemCount int
	head *sllNode
}

func InitSLL() sLL {
	return sLL{}
}


func (s *sLL) tranverse(_ bool) []int {
	var nums []int
	for tmp := s.head; tmp != nil; tmp = tmp.next {
		nums = append(nums, tmp.data)
	}
	return nums
}

func (s *sLL) search(searchData int) (int, bool) {
	for tmp, currIndex := s.head, 0; tmp != nil; tmp, currIndex = tmp.next, currIndex + 1 {
		if tmp.data == searchData {
			return currIndex, true
		}
	}
	return 0, false
}

func (s *sLL) insertion(inputData int, insertionIndex int) error {
	if insertionIndex < 0 {
		insertionIndex += s.itemCount + 1
	}
	if s.itemCount > 0 && insertionIndex > s.itemCount || insertionIndex < 0 {
		errorMsg := fmt.Sprintf("the minimum index of insertion is 0 while the maximum index is %d", s.itemCount)
		return errors.New(errorMsg)
	}

	newNode := sllNode {data: inputData}
	if insertionIndex == 0 {
		newNode.next = s.head
		s.head = &newNode
	} else {
		for tmp, currIndex := s.head, 0; tmp != nil; tmp, currIndex = tmp.next, currIndex + 1 {
			if currIndex == insertionIndex - 1 {
				newNode.next = tmp.next
				tmp.next = &newNode
				break
			}
		}
	}

	s.itemCount++
	return nil
}


func (s *sLL) retrieval(retrievalIndex int, isDelete bool) (int, error) {
	if retrievalIndex < 0 {
		retrievalIndex += s.itemCount
	}
	if s.itemCount > 0 && retrievalIndex >= s.itemCount || retrievalIndex < 0 {
		errorMsg := fmt.Sprintf("the minimum index of retrieval is 0 while the maximum index is %d", s.itemCount - 1)
		return 0, errors.New(errorMsg)
	}
	if s.itemCount == 0 {
		return 0, errors.New("the linked list is empty")
	}

	var val int
	if retrievalIndex == 0 {
		val = s.head.data
		if isDelete {
			s.head = s.head.next
		}
	} else {
		for tmp, currIndex := s.head, 0; tmp != nil; tmp, currIndex = tmp.next, currIndex + 1 {
			if currIndex == retrievalIndex - 1 {
				val = tmp.next.data
				if isDelete {
					tmp.next = tmp.next.next
				}
				break
			}
		}
	}

	if isDelete {	
		s.itemCount--
	}

	return val, nil
}
