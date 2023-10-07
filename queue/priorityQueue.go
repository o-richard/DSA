package queue

import "errors"

type priorityQueue struct {
	itemCount int
	maxSize int
	values []int
	isMaxHeap bool
}

func InitPriorityQueue(size int, isMaxHeap bool) (priorityQueue, error) {
	if size < 1 {
		return priorityQueue{}, errors.New("the size should have a minimum of 1")
	}

	var p priorityQueue
	p.maxSize = size
	p.isMaxHeap = isMaxHeap
	p.values = make([]int, size)
	return p, nil
}

func maxHeapify(arr []int, size, largestIndex int) {
	largest := largestIndex
	left := 2 * largestIndex + 1
	right := 2 * largestIndex + 2

	if left < size && arr[largestIndex] < arr[left] {
		largest = left
	}

	if right < size && arr[largest] < arr[right] {
		largest = right
	}

	if largest != largestIndex {
		arr[largestIndex], arr[largest] = arr[largest], arr[largestIndex]
		maxHeapify(arr, size, largest)
	}
}

func minHeapify(arr []int, size, smallestIndex int) {
	smallest := smallestIndex
	left := 2 * smallestIndex + 1
	right := 2 * smallestIndex + 2

	if left < size && arr[smallestIndex] > arr[left] {
		smallest = left
	}

	if right < size && arr[smallest] > arr[right] {
		smallest = right
	}

	if smallest != smallestIndex {
		arr[smallestIndex], arr[smallest] = arr[smallest], arr[smallestIndex]
		minHeapify(arr, size, smallest)
	}
	
}

func (p *priorityQueue) Insert(input int) error {
	if p.itemCount == p.maxSize {
		return errors.New("the queue is full")
	}

	if p.itemCount == 0 {
		p.values[0] = input
		p.itemCount++
	} else {
		p.values[p.itemCount] = input
		p.itemCount++
		for i := (p.itemCount / 2) - 1; i > -1; i-- {
			if p.isMaxHeap {
				maxHeapify(p.values, p.itemCount, i)
			} else {
				minHeapify(p.values, p.itemCount, i)
			}
		}
	}

	return nil
}

func (p *priorityQueue) Delete() (int, error) {
	if p.itemCount == 0 {
		return 0, errors.New("the queue is empty")
	}
	val := p.values[0]
	p.itemCount--
	p.values[0] = p.values[p.itemCount]
	p.values[p.itemCount] = 0
	for i := (p.itemCount / 2) - 1; i > -1; i-- {
		if p.isMaxHeap {
			maxHeapify(p.values, p.itemCount, i)
		} else {
			minHeapify(p.values, p.itemCount, i)
		}
	}
	return val, nil
}

func (p *priorityQueue) Peek() (int, error) {
	if p.itemCount == 0 {
		return 0, errors.New("the queue is empty")
	}

	return p.values[0], nil
}

func (p *priorityQueue) IsEmpty() bool {
	return p.itemCount == 0
}

func (p *priorityQueue) IsFull() bool {
	return p.itemCount == p.maxSize
}
