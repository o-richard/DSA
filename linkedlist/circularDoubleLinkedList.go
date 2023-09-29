package linkedlist

import (
	"fmt"
	"errors"
)

// cDllNode refers to a circularDoubleLinkedListNode
// cDLL refers to Circular Double Linked List

type cDllNode struct {
	data int
	prev *cDllNode
	next *cDllNode
}

type cDLL struct {
	itemCount int
	head *cDllNode
	tail *cDllNode
}

func InitCDLL() cDLL {
	return cDLL{}
}


func (c *cDLL) tranverse(goLeftToRight bool) []int {
	var nums []int
	if goLeftToRight {
		for tmp := c.head; tmp != nil; tmp = tmp.next {
			nums = append(nums, tmp.data)
			
			if tmp == c.tail {
				break
			}
		}
	} else {
		for tmp := c.tail; tmp != nil; tmp = tmp.prev {
			nums = append(nums, tmp.data)
			
			if tmp == c.head {
				break
			}
		}
	}
	
	return nums
}

func (c *cDLL) search(searchData int) (int, bool) {
	for tmp, currIndex := c.head, 0; tmp != nil; tmp, currIndex = tmp.next, currIndex + 1 {
		if tmp.data == searchData {
			return currIndex, true
		}

		if tmp == c.tail {
			break
		}
	}
	return 0, false
}

func (c *cDLL) insertion(inputData int, insertionIndex int) error {
	if insertionIndex < 0 {
		insertionIndex += c.itemCount + 1
	}
	if c.itemCount > 0 && insertionIndex > c.itemCount || insertionIndex < 0 {
		errorMsg := fmt.Sprintf("the minimum index of insertion is 0 while the maximum index is %d", c.itemCount)
		return errors.New(errorMsg)
	}

	newNode := cDllNode {data: inputData}
	if insertionIndex == 0 {
		newNode.next = c.head
		if c.itemCount > 0 {
			newNode.next.prev = &newNode
		}
		c.head = &newNode
		if c.tail == nil {
			c.tail = c.head
		}
		c.tail.next = c.head
		c.head.prev = c.tail
	} else {
		for tmp, currIndex := c.head, 0; ; tmp, currIndex = tmp.next, currIndex + 1 {
			if currIndex == insertionIndex - 1 {
				if tmp.next == c.head {
					c.tail = &newNode
				} 
				newNode.next = tmp.next
				newNode.next.prev = &newNode
				newNode.prev = tmp
				tmp.next = &newNode
				break
			}
		}
	}

	c.itemCount++
	return nil
}


func (c *cDLL) retrieval(retrievalIndex int, isDelete bool) (int, error) {
	if retrievalIndex < 0 {
		retrievalIndex += c.itemCount
	}
	if c.itemCount > 0 && retrievalIndex >= c.itemCount || retrievalIndex < 0 {
		errorMsg := fmt.Sprintf("the minimum index of retrieval is 0 while the maximum index is %d", c.itemCount - 1)
		return 0, errors.New(errorMsg)
	}
	if c.itemCount == 0 {
		return 0, errors.New("the linked list is empty")
	}

	var val int
	if retrievalIndex == 0 {
		val = c.head.data
		if isDelete {
			if c.itemCount == 1 {
				c.head = nil
				c.tail = nil
			} else {
				c.head = c.head.next
				c.head.prev = c.tail
				c.tail.next = c.head
			}
		}
	} else {
		for tmp, currIndex := c.head, 0; ; tmp, currIndex = tmp.next, currIndex + 1 {
			if currIndex == retrievalIndex - 1 {
				val = tmp.next.data
				if isDelete {
					tmp.next = tmp.next.next
					tmp.next.prev = tmp
					if tmp.next == c.head {
						c.tail = tmp
					}
				}
				break
			}
		}
	}

	if isDelete {	
		c.itemCount--
	}
	return val, nil
}
