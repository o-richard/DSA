package tree

import (
	"errors"
	"math"
)

type avlNode struct {
	data, height int
	left, right *avlNode
}

type avl struct {
	root *avlNode
}

func InitAVL() avl {
	return avl{}
}

func (n *avlNode) obtainHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *avlNode) getBalance() int {
	if n == nil {
		return 0
	}
	return n.left.obtainHeight() - n.right.obtainHeight()
}

func (n * avlNode) rightRotate() *avlNode {
	childLeft := n.left
	childLeftRight := n.left.right

	childLeft.right = n
	n.left = childLeftRight

	n.height = 1 + max(n.left.obtainHeight(), n.right.obtainHeight())
	childLeft.height = 1 + max(childLeft.left.obtainHeight(), childLeft.right.obtainHeight())

	return childLeft
}

func (n * avlNode) leftRotate() *avlNode {
	childRight := n.right
	childRightLeft := n.right.left

	childRight.left = n
	n.right = childRightLeft

	n.height = 1 + max(n.left.obtainHeight(), n.right.obtainHeight())
	childRight.height = 1 + max(childRight.left.obtainHeight(), childRight.right.obtainHeight())

	return childRight
}

func (n *avlNode) insertNode(inputData int) *avlNode {
	if inputData < n.data {
		if n.left == nil {
			n.left = &avlNode{data: inputData, height: 1}
		} else {
			n.left = n.left.insertNode(inputData)
		}
	} else {
		if n.right == nil {
			n.right = &avlNode{data: inputData, height: 1}
		} else {
			n.right = n.right.insertNode(inputData)
		}
	}

	n.height = 1 + max(n.left.obtainHeight(), n.right.obtainHeight())
	balanceFactor := n.getBalance()
	// The left subtree is unbalanced
	if balanceFactor > 1 {
		if inputData < n.left.data {
			return n.rightRotate()
		} else {
			n.left = n.left.leftRotate()
			return n.rightRotate()
		}
	}
	// The right subtree is unbalanced
	if balanceFactor < -1 {
		if inputData > n.right.data {
			return n.leftRotate()
		} else {
			n.right = n.right.rightRotate()
			return n.leftRotate()
		}
	}

	return n
}

func (a *avl) insert(inputData int) {
	if a.root == nil {
		a.root = &avlNode{data: inputData, height: 1}
	} else {
		a.root = a.root.insertNode(inputData)
	}
}

func (n *avlNode) searchNode(searchData int) bool {
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

func (a *avl) search(searchData int) bool {
	if a.root == nil {
		return false
	} else {
		return a.root.searchNode(searchData)
	}
}

func (n *avlNode) preOrder(nums []int) []int {
	if n != nil {
		nums = append(nums, n.data)
		nums = n.left.preOrder(nums)
		nums = n.right.preOrder(nums)
	}
	return nums
}

func (n *avlNode) inOrder(nums []int) []int {
	if n != nil {
		nums = n.left.inOrder(nums)
		nums = append(nums, n.data)
		nums = n.right.inOrder(nums)
	}
	return nums
}

func (n *avlNode) postOrder(nums []int) []int {
	if n != nil {
		nums = n.left.postOrder(nums)
		nums = n.right.postOrder(nums)
		nums = append(nums, n.data)
	}
	return nums
}

func (a *avl) tranverse(tranverseType string) ([]int, error) {
	var nums []int
	var err error

	switch tranverseType {
		case "preorder":
			nums = a.root.preOrder(nums)
			err = nil
		case "inorder":
			nums = a.root.inOrder(nums)
			err = nil
		case "postorder":
			nums = a.root.postOrder(nums)
			err = nil
		default:
			err = errors.New("unsupported tranverse type: options are preorder, postorder, inorder")	
	}
	return nums, err
}

func (n *avlNode) deleteNode(data int) (*avlNode, bool) {
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
			var tmp *avlNode
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

	if n == nil {
		return n, isPresent
	}

	n.height = 1 + max(n.left.obtainHeight(), n.right.obtainHeight())
	balanceFactor := n.getBalance()
	// The left subtree is unbalanced
	if balanceFactor > 1 {
		if n.left.getBalance() >= 0 {
			return n.rightRotate(), isPresent
		} else {
			n.left = n.left.leftRotate()
			return n.rightRotate(), isPresent
		}
	}
	// The right subtree is unbalanced
	if balanceFactor < -1 {
		if n.right.getBalance() <= 0 {
			return n.leftRotate(), isPresent
		} else {
			n.right = n.right.rightRotate()
			return n.leftRotate(), isPresent
		}
	}

	return n, isPresent
}

func (a *avl) delete(data int) error {
	if a.root == nil {
		return errors.New("the tree is empty")
	}
	var isPresent bool
	a.root, isPresent = a.root.deleteNode(data)
	if !isPresent {
		return errors.New("the node does not exist")
	}
	return nil
}

func (n *avlNode) isFullSubtree() bool {
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

func (a *avl) isFull() bool {
	return a.root.isFullSubtree()
}

func (n *avlNode) isPerfectSubtree(depth, level int) bool {
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

func (n *avlNode) calculateDepth() int {
	var depth int
	for n != nil {
		depth++
		n = n.left
	}
	return depth
}

func (a *avl) isPerfect() bool {
	depth, level := a.root.calculateDepth(), 0
	return a.root.isPerfectSubtree(depth, level)
}

func (n *avlNode) countNodes() int {
	if n == nil {
		return 0
	}
	return (1 + n.left.countNodes() + n.right.countNodes())
}

func (n *avlNode) isCompleteSubtree(index, nodeCount int) bool {
	if n == nil {
		return true
	}
	if index >= nodeCount {
		return false
	}
	return n.left.isCompleteSubtree(2 * index + 1, nodeCount) && n.right.isCompleteSubtree(2 * index + 2, nodeCount)
}

func (a *avl) isComplete() bool {
	index, nodeCount := 0, a.root.countNodes()
	return a.root.isCompleteSubtree(index, nodeCount)
}

func (n *avlNode) isBalancedSubTree(height *int) bool {
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

func (a *avl) isBalanced() bool {
	var height int
	return a.root.isBalancedSubTree(&height)
}
