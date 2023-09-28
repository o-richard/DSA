package linkedlist

import (
	"fmt"
	"errors"
)

// dllNode refers to a doubleLinkedListNode
// dLL refers to Double Linked List

type dllNode struct {
	data int
	prev *dllNode
	next *dllNode
}

type dLL struct {
	itemCount int
	head *dllNode
}

func InitDLL() dLL {
	return dLL{}
}

func (d *dLL) tranverse(_ bool) []int {
	var nums []int
	for tmp := d.head; tmp != nil; tmp = tmp.next {
		nums = append(nums, tmp.data)
	}
	return nums
}

func (d *dLL) search(searchData int) (int, bool) {
	for tmp, currIndex := d.head, 0; tmp != nil; tmp, currIndex = tmp.next, currIndex + 1 {
		if tmp.data == searchData {
			return currIndex, true
		}
	}
	return 0, false
}

func (d *dLL) insertion(inputData int, insertionIndex int) error {
	if insertionIndex < 0 {
		insertionIndex += d.itemCount + 1
	}
	if d.itemCount > 0 && insertionIndex > d.itemCount || insertionIndex < 0 {
		errorMsg := fmt.Sprintf("the minimum index of insertion is 0 while the maximum index is %d", d.itemCount)
		return errors.New(errorMsg)
	}

	newNode := dllNode {data: inputData}
	if insertionIndex == 0 {
		newNode.next = d.head
		if d.itemCount > 0 {
			newNode.next.prev = &newNode
		}
		d.head = &newNode
	} else {
		for tmp, currIndex := d.head, 0; tmp != nil; tmp, currIndex = tmp.next, currIndex + 1 {
			if currIndex == insertionIndex - 1 {
				newNode.next = tmp.next
				newNode.prev = tmp
				if newNode.next != nil {
					newNode.next.prev = &newNode 
				}
				tmp.next = &newNode
				break
			}
		}
	}

	d.itemCount++
	return nil
}

func (d *dLL) retrieval(retrievalIndex int, isDelete bool) (int, error) {
	if retrievalIndex < 0 {
		retrievalIndex += d.itemCount
	}
	if d.itemCount > 0 && retrievalIndex >= d.itemCount || retrievalIndex < 0 {
		errorMsg := fmt.Sprintf("the minimum index of retrieval is 0 while the maximum index is %d", d.itemCount - 1)
		return 0, errors.New(errorMsg)
	}
	if d.itemCount == 0 {
		return 0, errors.New("the linked list is empty")
	}

	var val int
	if retrievalIndex == 0 {
		val = d.head.data
		if isDelete {
			d.head = d.head.next
			if d.itemCount > 1 {
				d.head.prev = nil
			}
		}
	} else {
		for tmp, currIndex := d.head, 0; tmp != nil; tmp, currIndex = tmp.next, currIndex + 1 {
			if currIndex == retrievalIndex - 1 {
				val = tmp.next.data
				if isDelete {
					tmp.next = tmp.next.next
					if tmp.next != nil {
						tmp.next.prev = tmp
					}
				}
				break
			}
		}
	}

	if isDelete {	
		d.itemCount--
	}

	return val, nil
}
