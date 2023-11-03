package robinkarp_test

import (
	"reflect"
	"testing"

	"github.com/o-richard/DSA/robinkarp"
)

func TestRobinkarp(t *testing.T) {
	tests := map[string]struct {
		text, pattern   string
		expectedIndices []int
	}{
		"Test 1": {
			text:            "ABCCDDAEFG",
			pattern:         "CDD",
			expectedIndices: []int{3},
		},
		"Test 2": {
			text:            "AABAABAAB",
			pattern:         "AAB",
			expectedIndices: []int{0, 3, 6},
		},
		"Test 3": {
			text:            "AABAABAAB",
			pattern:         "XYZ",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := robinkarp.Robinkarp(test.text, test.pattern)
			if !reflect.DeepEqual(result, test.expectedIndices) {
				t.Errorf("expected the list of indices to be %v but got %v for the text %v and the pattern %v", test.expectedIndices, result, test.text, test.pattern)
			}
		})
	}
}
