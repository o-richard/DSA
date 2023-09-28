package tree

import (
	"errors"
	"math"
)

type bstNode struct {
	data int
	left, right *bstNode
}

type bst struct {
	root *bstNode
}

func InitBST() bst {
	return bst{}
}

func (n *bstNode) insertNode(inputData int) {
	if inputData < n.data {
		if n.left == nil {
			n.left = &bstNode{data: inputData}
			return
		} else {
			n.left.insertNode(inputData)
		}
	} else {
		if n.right == nil {
			n.right = &bstNode{data: inputData}
			return
		} else {
			n.right.insertNode(inputData)
		}
	}

}

func (b *bst) insert(inputData int) {
	if b.root == nil {
		b.root = &bstNode{data: inputData}
	} else {
		b.root.insertNode(inputData)
	}
}

func (n *bstNode) searchNode(searchData int) bool {
	var isFound bool
	if n != nil {	
		if n.data == searchData {
			isFound = true
		} else if searchData < n.data {
			isFound = n.left.searchNode(searchData)
		} else {
			isFound = n.right.searchNode(searchData)
		}
	}
	return isFound
}

func (b *bst) search(searchData int) bool {
	if b.root == nil {
		return false
	} else {
		return b.root.searchNode(searchData)
	}
}

func (n *bstNode) preOrder(nums []int) []int {
	if n != nil {
		nums = append(nums, n.data)
		nums = n.left.preOrder(nums)
		nums = n.right.preOrder(nums)
	}
	return nums
}

func (n *bstNode) inOrder(nums []int) []int {
	if n != nil {
		nums = n.left.inOrder(nums)
		nums = append(nums, n.data)
		nums = n.right.inOrder(nums)
	}
	return nums
}

func (n *bstNode) postOrder(nums []int) []int {
	if n != nil {
		nums = n.left.postOrder(nums)
		nums = n.right.postOrder(nums)
		nums = append(nums, n.data)
	}
	return nums
}

func (b *bst) tranverse(tranverseType string) ([]int, error) {
	var nums []int
	var err error

	switch tranverseType {
		case "preorder":
			nums = b.root.preOrder(nums)
			err = nil
		case "inorder":
			nums = b.root.inOrder(nums)
			err = nil
		case "postorder":
			nums = b.root.postOrder(nums)
			err = nil
		default:
			err = errors.New("unsupported tranverse type: options are preorder, postorder, inorder")	
	}
	return nums, err
}

func (n *bstNode) deleteNode(data int) (*bstNode, bool) {
	var isPresent bool
	if data < n.data {
		if n.left != nil {
			n.left, isPresent = n.left.deleteNode(data)
		}
	} else if data > n.data {
		if n.right != nil {
			n.right, isPresent = n.right.deleteNode(data)
		}
	} else {
		isPresent = true
		if n.left == nil && n.right == nil {
			n = nil
		} else if n.left != nil && n.right != nil {
			var tmp *bstNode
			for tmp = n.right; tmp.left != nil; tmp = tmp.left {
			}
			// Perform swap operation with the leftmost node of the current right node (in order successor)
			n.data, tmp.data = tmp.data, n.data
			n.right, isPresent = n.right.deleteNode(tmp.data)
		} else {
			if n.left != nil {
				n = n.left
			} else {
				n = n.right
			}
		}
	}
	return n, isPresent
}

func (b *bst) delete(data int) error {
	if b.root == nil {
		return errors.New("the tree is empty")
	}
	var isPresent bool
	b.root, isPresent = b.root.deleteNode(data)
	if !isPresent {
		return errors.New("the node does not exist")
	}
	return nil
}

func (n *bstNode) isFullSubtree() bool {
	if n == nil {
		return true
	}
	if n.left == nil && n.right == nil {
		return true
	}
	if n.left != nil && n.right != nil {
		return n.left.isFullSubtree() && n.right.isFullSubtree()
	}
	return false
}

func (b *bst) isFull() bool {
	return b.root.isFullSubtree()
}

func (n *bstNode) isPerfectSubtree(depth, level int) bool {
	if n == nil {
		return true
	}
	if n.left == nil && n.right == nil {
		return depth == level + 1
	}
	if n.left == nil || n.right == nil {
		return false
	}
	return n.left.isPerfectSubtree(depth, level + 1) && n.right.isPerfectSubtree(depth, level + 1)
}

func (n *bstNode) calculateDepth() int {
	var depth int
	for n != nil {
		depth++
		n = n.left
	}
	return depth
}

func (b *bst) isPerfect() bool {
	depth, level := b.root.calculateDepth(), 0
	return b.root.isPerfectSubtree(depth, level)
}

func (n *bstNode) countNodes() int {
	if n == nil {
		return 0
	}
	return (1 + n.left.countNodes() + n.right.countNodes())
}

func (n *bstNode) isCompleteSubtree(index, nodeCount int) bool {
	if n == nil {
		return true
	}
	if index >= nodeCount {
		return false
	}
	return n.left.isCompleteSubtree(2 * index + 1, nodeCount) && n.right.isCompleteSubtree(2 * index + 2, nodeCount)
}

func (b *bst) isComplete() bool {
	index, nodeCount := 0, b.root.countNodes()
	return b.root.isCompleteSubtree(index, nodeCount)
}

func (n *bstNode) isBalancedSubTree(height *int) bool {
	var leftHeight, rightHeght int
	
	if n == nil {
		*height = 0
		return true
	}

	l := n.left.isBalancedSubTree(&leftHeight)
	r := n.right.isBalancedSubTree(&rightHeght)

	*height = 1 + max(leftHeight, rightHeght)

	if math.Abs(float64(leftHeight - rightHeght)) <= 1 {
		return l && r
	} 
	return false
}

func (b *bst) isBalanced() bool {
	var height int
	return b.root.isBalancedSubTree(&height)
}
