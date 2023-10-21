package sort_test

import (
	"reflect"
	"testing"

	"github.com/o-richard/DSA/sort"
)


type Args struct {
	nums []int
	isAsc bool
}
type Expected struct {
	nums []int
	// Keeps track if an error will be obtained in the different sorting functions.
	// counting sort & radix sort is expected to yield an error incase of negative values.
	anyError bool
}
var tests = map[string]struct {
	agrs    Args
	expected Expected
}{
	"Test 1": {
		agrs: Args{
			nums: []int {-100, 1, 2, 3, 4},
			isAsc: true,
		},
		expected: Expected{
			nums: []int {-100, 1, 2, 3, 4},
			anyError: true,
		},
	},
	"Test 2": {
		agrs: Args{
			nums: []int {-100, 1, 2, 3, 4},
			isAsc: false,
		},
		expected: Expected{
			nums: []int {4, 3, 2, 1, -100},
			anyError: true,
		},
	},
	"Test 3": {
		agrs: Args{
			nums: []int {144, 9, 93, 9, 4, 7, 100},
			isAsc: true,
		},
		expected: Expected{
			nums: []int {4, 7, 9, 9, 93, 100, 144},
		},
	},
	"Test 4": {
		agrs: Args{
			nums: []int {100, 144, 9, 93, 9, 4, 7},
			isAsc: false,
		},
		expected: Expected{
			nums: []int {144, 100, 93, 9, 9, 7, 4},
		},
	},
	"Test 5": {
		agrs: Args{
			nums: []int {5, 2, 1, 2, 2},
			isAsc: true,
		},
		expected: Expected{
			nums: []int {1, 2, 2, 2, 5},
		},
	},
	"Test 6": {
		agrs: Args{
			nums: []int {5, 2, 1, 2, 2},
			isAsc: false,
		},
		expected: Expected{
			nums: []int {5, 2, 2, 2, 1},
		},
	},
}

func TestBubbleSort(t *testing.T) {
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            nums := sort.BubbleSort(test.agrs.nums, test.agrs.isAsc)
			if !reflect.DeepEqual(test.expected.nums, nums) {
				t.Errorf("For the array %v and isAsc %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.isAsc, test.expected.nums, nums)
			}
        })
    }
}

func TestCountingSort(t *testing.T) {
    for name, test := range tests {
		if !test.agrs.isAsc {
			continue
		}
        t.Run(name, func(t *testing.T) {
            nums, anyError := sort.CountingSort(test.agrs.nums)
			if anyError == nil && !reflect.DeepEqual(test.expected.nums, nums) {
				t.Errorf("For the array %v and isAsc %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.isAsc, test.expected.nums, nums)
			}
			if anyError == nil && test.expected.anyError {
				t.Errorf("For the array %v, an error is expected due to presence of negative elements", test.agrs.nums)
			}
        })
    }
}

func TestHeapSort(t *testing.T) {
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            nums := sort.HeapSort(test.agrs.nums, test.agrs.isAsc)
			if !reflect.DeepEqual(test.expected.nums, nums) {
				t.Errorf("For the array %v and isAsc %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.isAsc, test.expected.nums, nums)
			}
        })
    }
}

func TestInsertionSort(t *testing.T) {
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            nums := sort.InsertionSort(test.agrs.nums, test.agrs.isAsc)
			if !reflect.DeepEqual(test.expected.nums, nums) {
				t.Errorf("For the array %v and isAsc %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.isAsc, test.expected.nums, nums)
			}
        })
    }
}

func TestMergeSort(t *testing.T) {
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            nums := sort.MergeSort(test.agrs.nums, test.agrs.isAsc)
			if !reflect.DeepEqual(test.expected.nums, nums) {
				t.Errorf("For the array %v and isAsc %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.isAsc, test.expected.nums, nums)
			}
        })
    }
}

func TestQuickSort(t *testing.T) {
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            nums := sort.QuickSort(test.agrs.nums, test.agrs.isAsc)
			if !reflect.DeepEqual(test.expected.nums, nums) {
				t.Errorf("For the array %v and isAsc %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.isAsc, test.expected.nums, nums)
			}
        })
    }
}

func TestRadixSort(t *testing.T) {
    for name, test := range tests {
		if !test.agrs.isAsc {
			continue
		}
        t.Run(name, func(t *testing.T) {
            nums, anyError := sort.RadixSort(test.agrs.nums)
			if anyError == nil && !reflect.DeepEqual(test.expected.nums, nums) {
				t.Errorf("For the array %v and isAsc %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.isAsc, test.expected.nums, nums)
			}
			if anyError == nil && test.expected.anyError {
				t.Errorf("For the array %v, an error is expected due to presence of negative elements", test.agrs.nums)
			}
        })
    }
}

func TestSelectionSort(t *testing.T) {
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            nums := sort.SelectionSort(test.agrs.nums, test.agrs.isAsc)
			if !reflect.DeepEqual(test.expected.nums, nums) {
				t.Errorf("For the array %v and isAsc %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.isAsc, test.expected.nums, nums)
			}
        })
    }
}

func TestShellSort(t *testing.T) {
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            nums := sort.ShellSort(test.agrs.nums, test.agrs.isAsc)
			if !reflect.DeepEqual(test.expected.nums, nums) {
				t.Errorf("For the array %v and isAsc %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.isAsc, test.expected.nums, nums)
			}
        })
    }
}