package tree_test

import (
	"reflect"
	"testing"

	"github.com/o-richard/DSA/tree"
)

var nums = []int{7,8,9,1,2,3,4,5,7,8,8,8,10,11, -100}

func TestBST(t *testing.T) {
	mytree := tree.InitBST()
	
	for _, num := range nums {
		tree.Insert(&mytree, num)
	}
	for _, num := range nums {
		err := tree.Delete(&mytree, num)
		if err != nil {
			t.Fatalf("deletion of all elements in the tree was not successful. Either insertion or deletion is not successful.")
		}
	}

	for _, num := range nums {
		tree.Insert(&mytree, num)
	}
	for _, num := range nums {
		isFound := tree.Search(&mytree, num)
		if !isFound {
			t.Errorf("failed to find %v in the tree", num)
		}
	}
	isFound := tree.Search(&mytree, 44)
	if isFound {
		t.Errorf("found %v and it is absent in the tree", 44)
	}

	_, err:= tree.Tranverse(&mytree, "test")
	if err == nil {
		t.Errorf("test should be an unsupported tranverse type. Expected options are preorder, inorder, postorder")
	}

	// the output should be expected since insertion into the tree is 'controlled', no deletion has been done.
	tests := map[string]struct {
		order string
		expectedNums []int
	}{
		"Test 1": {
			order: "preorder",
			expectedNums: []int {7, 1, -100, 2, 3, 4, 5, 8, 7, 9, 8, 8, 8, 10, 11},
		},
		"Test 2": {
			order: "inorder",
			expectedNums: []int {-100, 1, 2, 3, 4, 5, 7, 7, 8, 8, 8, 8, 9, 10, 11},
		},
		"Test 3": {
			order: "postorder",
			expectedNums: []int {-100, 5, 4, 3, 2, 1, 7, 8, 8, 8, 11, 10, 9, 8, 7},
		},
	}
	for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual, err := tree.Tranverse(&mytree, test.order)
			if !reflect.DeepEqual(actual, test.expectedNums) {
				t.Errorf("expected %v while performing %v tranverse type, actual result is %v", test.expectedNums, test.order, actual)
			}
			if err != nil {
				t.Errorf("an error was unexpected for %v tranverse type", test.order)
			}
        })
    }
}

func TestAVL(t *testing.T) {
	mytree := tree.InitAVL()
	
	for _, num := range nums {
		tree.Insert(&mytree, num)
	}
	for _, num := range nums {
		err := tree.Delete(&mytree, num)
		if err != nil {
			t.Fatalf("deletion of all elements in the tree was not successful. Either insertion or deletion is not successful.")
		}
	}

	for _, num := range nums {
		tree.Insert(&mytree, num)
	}
	for _, num := range nums {
		isFound := tree.Search(&mytree, num)
		if !isFound {
			t.Errorf("failed to find %v in the tree", num)
		}
	}
	isFound := tree.Search(&mytree, 44)
	if isFound {
		t.Errorf("found %v and it is absent in the tree", 44)
	}

	isBalanced := tree.IsBalanced(&mytree)
	if !isBalanced {
		t.Errorf("the tree should be balanced")
	}

	_, err:= tree.Tranverse(&mytree, "test")
	if err == nil {
		t.Errorf("test should be an unsupported tranverse type. Expected options are preorder, inorder, postorder")
	}

	tests := map[string]struct {
		order string
		expectedNums []int
	}{
		"Test 1": {
			order: "preorder",
			expectedNums: []int {7, 2, 1, -100, 4, 3, 5, 8, 8, 7, 9, 8, 8, 10, 11},
		},
		"Test 2": {
			order: "inorder",
			expectedNums: []int {-100, 1, 2, 3, 4, 5, 7, 7, 8, 8, 8, 8, 9, 10, 11},
		},
		"Test 3": {
			order: "postorder",
			expectedNums: []int {-100, 1, 3, 5, 4, 2, 7, 8, 8, 8, 11, 10, 9, 8, 7},
		},
	}
	for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual, err := tree.Tranverse(&mytree, test.order)
			if !reflect.DeepEqual(actual, test.expectedNums) {
				t.Errorf("expected %v while performing %v tranverse type, actual result is %v", test.expectedNums, test.order, actual)
			}
			if err != nil {
				t.Errorf("an error was unexpected for %v tranverse type", test.order)
			}
        })
    }	
}