package heap_test

import (
	"fmt"
	"testing"

	"github.com/o-richard/DSA/heap"
)

func TestFHeap_FindRootFHeap(t *testing.T) {
	tests := map[string]struct {
		isMaxHeap       bool
		inputData       []int
		expectedRoot    int
		expectedSuccess bool
	}{
		"Max Fibonacci Heap":       {isMaxHeap: true, inputData: []int{6, 8, 1}, expectedRoot: 8, expectedSuccess: true},
		"Min Fibonacci Heap":       {inputData: []int{6, 8, 1}, expectedRoot: 1, expectedSuccess: true},
		"Empty Max Fibonacci Heap": {isMaxHeap: true, expectedRoot: 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			fheap := heap.InitFHeap(test.isMaxHeap)
			for _, data := range test.inputData {
				fheap.InsertFHeap(data)
			}
			actualRoot, err := fheap.FindRootFHeap()
			if test.expectedSuccess && test.expectedRoot != actualRoot {
				t.Errorf("expected the value %v but got %v for the input data %v", test.expectedRoot, actualRoot, test.inputData)
			}
			if !test.expectedSuccess && (err == nil) {
				t.Errorf("expected an error for an empty fibonacci heap")
			}
		})
	}
}

func TestFHeap_UnionFHeap(t *testing.T) {
	fheap1 := heap.InitFHeap(false)
	fheap1.InsertFHeap(3)
	fheap1.InsertFHeap(2)
	fheap1.InsertFHeap(1)

	fheap2 := heap.InitFHeap(false)
	fheap2.InsertFHeap(6)
	fheap2.InsertFHeap(7)
	fheap2.InsertFHeap(8)

	fheap3 := heap.InitFHeap(false)

	err := fheap1.UnionFHeap(&fheap2)
	if err != nil {
		t.Errorf("expected no error when unifying two fibonacci heaps")
	}
	err = fheap1.UnionFHeap(&fheap3)
	if err == nil {
		t.Errorf("expected an error when unifying a fibonacci heap with an empty one")
	}
}

func TestFHeap_ExtractRootFHeap(t *testing.T) {
	data := []int{6, 8, 1, 3}
	fheap := heap.InitFHeap(false)
	for _, v := range data {
		fheap.InsertFHeap(v)
	}

	tests := []struct {
		expectedRoot    int
		expectedSuccess bool
	}{
		{expectedRoot: 1, expectedSuccess: true},
		{expectedRoot: 3, expectedSuccess: true},
		{expectedRoot: 6, expectedSuccess: true},
		{expectedRoot: 8, expectedSuccess: true},
		{expectedRoot: 0},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			actualRoot, err := fheap.ExtractRootFHeap()
			if test.expectedSuccess && actualRoot != test.expectedRoot {
				t.Errorf("expected %v but got %v for the input %v", test.expectedRoot, actualRoot, data)
			}
			if !test.expectedSuccess && err == nil {
				t.Errorf("expected an error since the fibonacci heap should be empty")
			}
		})
	}
}

func TestFHeap_ChangeKey(t *testing.T) {
	data := []int{6, 8, 1, 3}
	fheap := heap.InitFHeap(false)
	for _, v := range data {
		fheap.InsertFHeap(v)
	}

	tests := map[string]struct {
		currentData     int
		newData         int
		expectedSuccess bool
	}{
		"Present Key": {currentData: 6, newData: 10, expectedSuccess: true},
		"Absent Key":  {currentData: 9, newData: 10},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := fheap.ChangeKey(test.currentData, test.newData)
			if test.expectedSuccess && err != nil {
				t.Errorf("expected no error since the item is present in the fibonacci heap")
			}
			if !test.expectedSuccess && err == nil {
				t.Errorf("expected an error since the item is absent in the fibonacci heap")
			}
		})
	}
}

func TestFHeap_DeleteNode(t *testing.T) {
	data := []int{6, 8, 1, 3}
	fheap := heap.InitFHeap(false)
	for _, v := range data {
		fheap.InsertFHeap(v)
	}

	tests := map[string]struct {
		currentData     int
		expectedSuccess bool
	}{
		"Present Key": {currentData: 6, expectedSuccess: true},
		"Absent Key":  {currentData: 9},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := fheap.DeleteNode(test.currentData)
			if test.expectedSuccess && err != nil {
				t.Errorf("expected no error since the item is present in the fibonacci heap")
			}
			if !test.expectedSuccess && err == nil {
				t.Errorf("expected an error since the item is absent in the fibonacci heap")
			}
		})
	}
}

func TestFHeap_UnionFibonacciHeap(t *testing.T) {
	fheap := heap.InitFHeap(false)
	fheap.InsertFHeap(3)
	fheap.InsertFHeap(2)
	fheap.InsertFHeap(1)

	fheap2 := heap.InitFHeap(false)
	fheap2.InsertFHeap(6)
	fheap2.InsertFHeap(7)
	fheap2.InsertFHeap(8)

	fheap.UnionFHeap(&fheap2)

	data := []int{6, 7, 8, 3, 1, 2}
	for _, currentData := range data {
		err := fheap.DeleteNode(currentData)
		if err != nil {
			t.Errorf("expected no error since the item is present in the fibonacci heap. There is an error likely during the union of two fibonacci heaps")
		}
	}
}
