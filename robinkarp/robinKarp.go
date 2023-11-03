package robinkarp

import (
	"strings"
)


func Robinkarp(text, pattern string) []int {
	pattern = strings.ToUpper(pattern)
	text = strings.ToUpper(text)

	patternLength := len(pattern)
	textLength := len(text)
	
	multiplier := 256
	mod := 101

	// Should match (multiplier ^ (patternLength - 1) % mod)
	maxHashConst := 1
	for i := 0; i < patternLength - 1; i++ {
		maxHashConst = (maxHashConst * multiplier) % mod
	}

	var patternHash, textHash int
	for i := 0; i < patternLength; i++ {
		patternHash = (patternHash * multiplier + int(pattern[i])) % mod
		textHash = (textHash * multiplier + int(text[i])) % mod
	}

	var indices []int

	for i := 0; i < textLength - patternLength + 1; i++ {
		if patternHash == textHash {
			isFound := true
			for j := 0; j < patternLength; j++ {
				if pattern[j] != text[i + j] {
					isFound = false
					break
				}
			}
			if isFound {
				indices = append(indices, i)
			}
		}
		if i < textLength - patternLength {
			textHash = ((textHash - int(text[i]) * maxHashConst) * multiplier + int(text[i+patternLength])) % mod
			if textHash < 0 {
				textHash += mod
			}	
		}
	}

	return indices
}