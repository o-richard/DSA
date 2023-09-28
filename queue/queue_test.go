package queue_test

import (
	"fmt"
	"testing"

	"github.com/o-richard/DSA/queue"
)

var initializationTests = map[string]struct {
	size            int
	expectedSuccess bool
}{
	"Valid Positive Size":   {size: 3, expectedSuccess: true},
	"Invalid Negative Size": {size: -3},
}

func TestCircularQueue_Init(t *testing.T) {
	for name, test := range initializationTests {
		t.Run(name, func(t *testing.T) {
			_, err := queue.InitCircularQueue(test.size)
			if test.expectedSuccess && err != nil {
				t.Errorf("expected no error when creating a circular queue of size %v", test.size)
			}
			if !test.expectedSuccess && err == nil {
				t.Errorf("expected an error when creating a circular queue of negative size %v", test.size)
			}
		})
	}
}

func TestCircularQueue_Enqueue(t *testing.T) {
	cq, _ := queue.InitCircularQueue(1)
	tests := []struct {
		data            int
		expectedSuccess bool
	}{
		{data: 1, expectedSuccess: true},
		{data: -3},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			err := cq.Enqueue(test.data)
			if test.expectedSuccess && err != nil {
				t.Errorf("expected no error during the enqueue operation of the circular queue")
			}
			if !test.expectedSuccess && err == nil {
				t.Errorf("expected an error during the enqueue operation of the circular queue. The queue should be full.")
			}
		})
	}
}

func TestCircularQueue(t *testing.T) {
	cq, _ := queue.InitCircularQueue(4)

	// IsEmpty Test
	if !cq.IsEmpty() {
		t.Errorf("expected the queue to be empty")
	}

	data := []int{2, 3, 4, 5}
	for _, v := range data {
		cq.Enqueue(v)
	}

	// Peek Test
	expectedPeek := 2
	actualPeek := cq.Peek()
	if actualPeek != expectedPeek {
		t.Errorf("expected %v at the front of the queue but got %v", expectedPeek, actualPeek)
	}

	// IsFull Test
	if !cq.IsFull() {
		t.Errorf("expected the queue to be full")
	}

	// Dequeue Test
	tests := []int{2, 3, 4, 5}

	for i, expectedData := range tests {
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			actualData := cq.Dequeue()
			if actualData != expectedData {
				t.Errorf("expected %v but got %v during the dequeue operation", expectedData, actualData)
			}
		})
	}
}

func TestDoubleEndedQueue_Init(t *testing.T) {
	for name, test := range initializationTests {
		t.Run(name, func(t *testing.T) {
			_, err := queue.InitDoubleEndedQueue(test.size)
			if test.expectedSuccess && err != nil {
				t.Errorf("expected no error when creating a double ended queue of size %v", test.size)
			}
			if !test.expectedSuccess && err == nil {
				t.Errorf("expected an error when creating a double ended queue of negative size %v", test.size)
			}
		})
	}
}

func TestDoubleEndedQueue_Add(t *testing.T) {
	dq, _ := queue.InitDoubleEndedQueue(5)

	tests := []struct {
		isFront         bool
		data            int
		expectedSuccess bool
	}{
		{isFront: true, data: 1, expectedSuccess: true},
		{isFront: true, data: 2, expectedSuccess: true},
		{isFront: true, data: 3, expectedSuccess: true},
		{isFront: false, data: 4, expectedSuccess: true},
		{isFront: false, data: 5, expectedSuccess: true},
		{isFront: true, data: 6},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			var err error
			if test.isFront {
				err = dq.AddAtFront(test.data)
			} else {
				err = dq.AddAtRear(test.data)
			}
			addingPosition := "front"
			if !test.isFront {
				addingPosition = "rear"
			}
			if test.expectedSuccess && err != nil {
				t.Errorf("expected no error when adding to the %v of the double ended queue", addingPosition)
			}
			if !test.expectedSuccess && err == nil {
				t.Errorf("expected an error when adding to the %v of a double ended queue. The queue should be full", addingPosition)
			}
		})
	}
}

func TestDoubleEndedQueue(t *testing.T) {
	dq, _ := queue.InitDoubleEndedQueue(5)

	// IsEmpty Test
	if !dq.IsEmpty() {
		t.Errorf("expected the double ended queue to be empty")
	}

	insertionData := []struct {
		isFront bool
		data    int
	}{
		{isFront: true, data: 1},
		{isFront: false, data: 4},
		{isFront: true, data: 2},
		{isFront: false, data: 5},
		{isFront: true, data: 3},
	}
	for _, v := range insertionData {
		if v.isFront {
			dq.AddAtFront(v.data)
		} else {
			dq.AddAtRear(v.data)
		}
	}

	// IsFull Test
	if !dq.IsFull() {
		t.Errorf("expected the double ended queue to be full")
	}

	// Delete Test
	tests := []struct {
		isFront bool
		data    int
	}{
		{isFront: false, data: 5},
		{isFront: true, data: 3},
		{isFront: true, data: 2},
		{isFront: false, data: 4},
		{isFront: true, data: 1},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			var actualData interface{}
			if test.isFront {
				actualData = dq.DeleteAtFront()
			} else {
				actualData = dq.DeleteAtRear()
			}
			deletionPosition := "front"
			if !test.isFront {
				deletionPosition = "rear"
			}
			if actualData != test.data {
				t.Errorf("expected %v when deleting from the %v of the double ended queue but got %v", test.data, deletionPosition, actualData)
			}
		})
	}
}

func TestPriorityQueue_Init(t *testing.T) {
	for name, test := range initializationTests {
		t.Run(name, func(t *testing.T) {
			_, err := queue.InitPriorityQueue(test.size, false)
			if test.expectedSuccess && err != nil {
				t.Errorf("expected no error when creating a priority queue of size %v", test.size)
			}
			if !test.expectedSuccess && err == nil {
				t.Errorf("expected an error when creating a priority queue of negative size %v", test.size)
			}
		})
	}
}

func TestPriorityQueue_Insert(t *testing.T) {
	pq, _ := queue.InitPriorityQueue(1, true)
	tests := []struct {
		data            int
		expectedSuccess bool
	}{
		{
			data:            1,
			expectedSuccess: true,
		},
		{
			data: -3,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			err := pq.Insert(test.data)
			if test.expectedSuccess && err != nil {
				t.Errorf("expected no error during the insertion operation of the priority queue")
			}
			if !test.expectedSuccess && err == nil {
				t.Errorf("expected an error during the insertion operation of the priority queue. The queue should be full.")
			}
		})
	}
}

func TestPriorityQueue(t *testing.T) {
	tests := map[string]struct {
		heapSize      int
		isMaxHeap     bool
		insertionData []int
		removalCount  int
		updateData    []int
		expectedOrder []int
	}{
		"Test 1": {
			heapSize:      4,
			isMaxHeap:     true,
			insertionData: []int{1, 2, 3, 4},
			removalCount:  2,
			updateData:    []int{6, 7},
			expectedOrder: []int{7, 6, 2, 1},
		},
		"Test 2": {
			heapSize:      4,
			isMaxHeap:     false,
			insertionData: []int{1, 2, 3, 4},
			removalCount:  2,
			updateData:    []int{6, 7},
			expectedOrder: []int{3, 4, 6, 7},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pq, _ := queue.InitPriorityQueue(test.heapSize, test.isMaxHeap)

			// IsEmpty test
			if !pq.IsEmpty() {
				t.Errorf("expected the priority queue to be empty")
			}

			for _, v := range test.insertionData {
				pq.Insert(v)
			}

			// IsFull test
			if !pq.IsFull() {
				t.Errorf("expected the priority queue to be empty")
			}

			for i := 0; i < test.removalCount; i++ {
				pq.Delete()
			}
			for _, v := range test.updateData {
				pq.Insert(v)
			}

			// Peek Test
			actualPeek, _ := pq.Peek()
			if actualPeek != test.expectedOrder[0] {
				t.Errorf("expected %v at the start of the priority queue but got %v", test.expectedOrder[0], actualPeek)
			}

			for _, v := range test.expectedOrder {
				actualOutput, _ := pq.Delete()
				if v != actualOutput {
					t.Errorf("expected %v while performing the delete operation in the priority queue but got %v", v, actualOutput)
				}
			}
		})
	}
}
