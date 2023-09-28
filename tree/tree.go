package tree

type treeType interface {
	insert(data int)
	search(data int) bool
	tranverse(tranverseType string) ([]int, error)
	delete(data int) error
	isFull() bool
	isPerfect() bool
	isComplete() bool
	isBalanced() bool
}

func Insert(tree treeType, data int) {
	tree.insert(data)
}

func Search(tree treeType, data int) bool {
	return tree.search(data)
}

// tranverseType options are preorder, inorder, postorder
func Tranverse(tree treeType, tranverseType string) ([]int, error) {
	return tree.tranverse(tranverseType)
}

func Delete(tree treeType, data int) error {
	return tree.delete(data)
}

func IsFull(tree treeType) bool {
	return tree.isFull()
}

func IsPerfect(tree treeType) bool {
	return tree.isPerfect()
}

func IsComplete(tree treeType) bool {
	return tree.isComplete()
}

func IsBalanced(tree treeType) bool {
	return tree.isBalanced()
}
