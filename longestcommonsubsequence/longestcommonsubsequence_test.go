package longestcommonsubsequence_test

import (
	"testing"

	"github.com/o-richard/DSA/longestcommonsubsequence"
)

func TestLCS(t *testing.T) {
	tests := map[string]struct {
		s1, s2        string
		expectedValue string
	}{
		"Test 1": {s1: "ABCBDAB", s2: "BDABC", expectedValue: "BDAB"},
		"Test 2": {s1: "BDABC", s2: "BDABC", expectedValue: "BDABC"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualOutput := longestcommonsubsequence.LCS(test.s1, test.s2)
			if actualOutput != test.expectedValue {
				t.Errorf("expected %v but got %v for the input %v and %v", test.expectedValue, actualOutput, test.s1, test.s2)
			}
		})
	}

}
