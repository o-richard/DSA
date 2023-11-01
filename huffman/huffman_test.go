package huffman_test

import (
	"reflect"
	"testing"

	"github.com/o-richard/DSA/huffman"
)

func TestHuffman(t *testing.T) {
	inputData := "BCAADDDCCACACAC"
	expectedOutput := map[string]string{
		"A": "11",
		"B": "100",
		"C": "0",
		"D": "101",
	}

	actual := huffman.Huffman(inputData)
	if !reflect.DeepEqual(actual, expectedOutput) {
		t.Errorf("expected %v but got %v for the input %v", expectedOutput, actual, inputData)
	}
}
