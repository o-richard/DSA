package search_test

import (
	"testing"

	"github.com/o-richard/DSA/search"
)

type Args struct {
	nums []int
	searchData int
}
type Expected struct {
	isFound bool
}
var tests = map[string]struct {
	agrs    Args
	expected Expected
}{
	"Test 1": {
		agrs: Args{
			nums: []int {1, 2, 3, 4},
			searchData: 5,
		},
		expected: Expected{
			isFound: false,
		},
	},
	"Test 2": {
		agrs: Args{
			nums: []int {1, 2, 3, 4},
			searchData: 2,
		},
		expected: Expected{
			isFound: true,
		},
	},
	"Test 3": {
		agrs: Args{
			nums: []int {},
			searchData: 2,
		},
		expected: Expected{
			isFound: false,
		},
	},
}

func TestLinearSearch(t *testing.T) {
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            isFound := search.LinearSearch(test.agrs.nums, test.agrs.searchData)
			if isFound != test.expected.isFound{
				t.Errorf("For the array %v and search data %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.searchData, test.expected.isFound, isFound)
			}
        })
    }
}

func TestBinarySearch(t *testing.T) {
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            isFound := search.BinarySearch(test.agrs.nums, test.agrs.searchData)
			if isFound != test.expected.isFound{
				t.Errorf("For the array %v and search data %v, the expected ouput is (%v). Actual output is (%v)", test.agrs.nums, test.agrs.searchData, test.expected.isFound, isFound)
			}
        })
    }
}