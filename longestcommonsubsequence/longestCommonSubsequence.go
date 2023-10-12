package longestcommonsubsequence

import "strings"

// LCS - Longest Common Subsequence

func LCS(s1, s2 string) string {
	s1 = strings.ToUpper(s1)
	s2 = strings.ToUpper(s2)
	
	s1Length := len(s1)
	s2Length := len(s2)

	result := make([][]int, s1Length + 1)
	for i := range result {
		result[i] = make([]int, s2Length + 1)
	}

	for i := 1; i <= s1Length; i++ {
		for j := 1; j <= s2Length; j++ {
			if s1[i - 1] == s2[j - 1] {
				result[i][j] = result[i - 1][j - 1] + 1
			} else {
				result[i][j] = max(result[i - 1][j], result[i][j - 1])
			}
		}
	}

	index := result[s1Length][s2Length]
	lcs := make([]byte, index)

	for i, j := s1Length, s2Length; i > 0 && j > 0; {
		if s1[i - 1] == s2[j - 1] {
			lcs[index - 1] = s1[i - 1]
			i--
			j--
			index--
		} else if result[i - 1][j] > result[i][j - 1] {
			i--
		} else {
			j--
		}
	}

	return string(lcs)
}
