package linkedlist_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/o-richard/DSA/linkedlist"
)

func TestInsertion(t *testing.T) {
	sll := linkedlist.InitSLL()
	dll := linkedlist.InitDLL()
	csll := linkedlist.InitCSLL()
	cdll := linkedlist.InitCDLL()
	ll := map[string]linkedlist.LinkedList{
		"Single Linked List":          &sll,
		"Double Linked List":          &dll,
		"Circular Single Linked List": &csll,
		"Circular Double Linked List": &cdll,
	}

	tests := []struct {
		insertionIndex  int
		inputData       int
		expectedSuccess bool
	}{
		{
			insertionIndex:  0,
			inputData:       1,
			expectedSuccess: true,
		},
		{
			insertionIndex:  -1,
			inputData:       2,
			expectedSuccess: true,
		},
		{
			insertionIndex:  1,
			inputData:       3,
			expectedSuccess: true,
		},
		{
			insertionIndex: -7,
			inputData:      4,
		},
	}

	for llName, llType := range ll {
		for i, test := range tests {
			t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
				err := linkedlist.Insertion(llType, test.inputData, test.insertionIndex)
				if test.expectedSuccess && err != nil {
					t.Errorf("%v: expected no error during insertion since provided indices are in the range of the current data in the linked list", llName)
				}
				if !test.expectedSuccess && err == nil {
					t.Errorf("%v: expected an error during insertion since provided indices are not in the range of the current data in the linked list", llName)
				}
			})
		}
	}
}

func TestSearch(t *testing.T) {
	sll := linkedlist.InitSLL()
	dll := linkedlist.InitDLL()
	csll := linkedlist.InitCSLL()
	cdll := linkedlist.InitCDLL()
	ll := map[string]linkedlist.LinkedList{
		"Single Linked List":          &sll,
		"Double Linked List":          &dll,
		"Circular Single Linked List": &csll,
		"Circular Double Linked List": &cdll,
	}

	inputData := []int{1, 3, 4, 6, 8}

	tests := map[string]struct {
		searchData     int
		expectedIndex  int
		expectedAnswer bool
	}{
		"Test 1": {
			searchData:     1,
			expectedIndex:  4,
			expectedAnswer: true,
		},
		"Test 2": {
			searchData:     6,
			expectedIndex:  1,
			expectedAnswer: true,
		},
		"Test 3": {
			searchData:     8,
			expectedAnswer: true,
		},
		"Test 4": {
			searchData: 28,
		},
	}

	for llName, llType := range ll {
		for _, v := range inputData {
			linkedlist.Insertion(llType, v, 0)
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				actualIndex, isFound := linkedlist.Search(llType, test.searchData)
				if test.expectedAnswer && actualIndex != test.expectedIndex || isFound != test.expectedAnswer {
					t.Errorf("%v: expected to find the value %v at index %v but got the index %v ", llName, test.searchData, test.expectedIndex, actualIndex)
				}
				if !test.expectedAnswer && actualIndex != test.expectedIndex || isFound != test.expectedAnswer {
					t.Errorf("%v: expected the absence of the value %v but got the index %v ", llName, test.searchData, actualIndex)
				}
			})
		}
	}
}

func TestRetrieval(t *testing.T) {
	sll := linkedlist.InitSLL()
	dll := linkedlist.InitDLL()
	csll := linkedlist.InitCSLL()
	cdll := linkedlist.InitCDLL()
	ll := map[string]linkedlist.LinkedList{
		"Single Linked List":          &sll,
		"Double Linked List":          &dll,
		"Circular Single Linked List": &csll,
		"Circular Double Linked List": &cdll,
	}

	inputData := []int{1, 3, 4, 6, 8}

	// Iteration of Maps tend to be randomized
	// The insertion order needs to be preserved for the tests
	tests := []struct {
		retrievalIndex  int
		isDelete        bool
		expectedValue   int
		expectedSuccess bool
	}{
		{
			retrievalIndex: 5,
		},
		{
			retrievalIndex:  -3,
			expectedValue:   4,
			expectedSuccess: true,
		},
		{
			retrievalIndex:  0,
			isDelete:        true,
			expectedValue:   8,
			expectedSuccess: true,
		},
		{
			retrievalIndex: 4,
		},
		{
			retrievalIndex:  0,
			expectedValue:   6,
			expectedSuccess: true,
		},
	}

	for llName, llType := range ll {
		for _, v := range inputData {
			linkedlist.Insertion(llType, v, 0)
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
				actualValue, err := linkedlist.Retrieval(llType, test.retrievalIndex, test.isDelete)
				if test.expectedSuccess && err != nil || actualValue != test.expectedValue {
					t.Errorf("%v: expected the value %v but got %v. No error is expected during retrieval since provided indices are in the range of the current data in the linked list", llName, test.expectedValue, actualValue)
				}
				if !test.expectedSuccess && err == nil {
					t.Errorf("%v: expected an error during retrieval since provided indices are not in the range of the current data in the linked list", llName)
				}
			})
		}
	}

}

func TestTranverse(t *testing.T) {
	sll := linkedlist.InitSLL()
	dll := linkedlist.InitDLL()
	csll := linkedlist.InitCSLL()
	cdll := linkedlist.InitCDLL()
	ll := map[string]linkedlist.LinkedList{
		"Single Linked List":          &sll,
		"Double Linked List":          &dll,
		"Circular Single Linked List": &csll,
		"Circular Double Linked List": &cdll,
	}

	inputData := []struct {
		insertionIndex int
		inputData      int
	}{
		{insertionIndex: 0, inputData: 1},
		{insertionIndex: 0, inputData: 3},
		{insertionIndex: -1, inputData: 4},
		{insertionIndex: -1, inputData: 6},
		{insertionIndex: 2, inputData: 8},
		{insertionIndex: 4, inputData: 9},
		{insertionIndex: 1, inputData: 11},
		{insertionIndex: 0, inputData: 22},
		{insertionIndex: -1, inputData: 44},
	}
	removalData := []int{0, -1}

	tests := map[string]struct {
		goLeftToRight  bool
		expectedOutput []int
	}{
		"Test 1": {
			goLeftToRight:  true,
			expectedOutput: []int{3, 11, 1, 8, 4, 9, 6},
		},
		"Test 2": {
			expectedOutput: []int{6, 9, 4, 8, 1, 11, 3},
		},
	}

	for llName, llType := range ll {
		for _, v := range inputData {
			linkedlist.Insertion(llType, v.inputData, v.insertionIndex)
		}
		for _, v := range removalData {
			linkedlist.Retrieval(llType, v, true)
		}
		for name, test := range tests {
			if !test.goLeftToRight && llName != "Circular Double Linked List" {
				continue
			}
			t.Run(name, func(t *testing.T) {
				actualOutput := linkedlist.Tranverse(llType, test.goLeftToRight)
				if !reflect.DeepEqual(actualOutput, test.expectedOutput) {
					t.Errorf("%v: expected %v but got %v during tranversal", llName, test.expectedOutput, actualOutput)
				}
			})
		}
	}

}
