package tree

var Templates = map[string]string{
	"AVL": `package tree

type avl struct {
	// Edit here
}

func InitAVL() avl {
	// Edit here
}

func (a *avl) insert(inputData int) {
	// Edit here
}

func (a *avl) search(searchData int) bool {
	// Edit here
}

func (a *avl) tranverse(tranverseType string) ([]int, error) {
	// Edit here
}

func (a *avl) delete(data int) error {
	// Edit here
}

func (a *avl) isFull() bool {
	// Edit here
}

func (a *avl) isPerfect() bool {
	// Edit here
}

func (a *avl) isComplete() bool {
	// Edit here
}

func (a *avl) isBalanced() bool {
	// Edit here	
}	
	`,
	"BST": `package tree

type bst struct {
	// Edit here
}

func InitBST() bst {
	// Edit here
}

func (b *bst) insert(inputData int) {
	// Edit here
}

func (b *bst) search(searchData int) bool {
	// Edit here
}

func (b *bst) tranverse(tranverseType string) ([]int, error) {
	// Edit here
}

func (b *bst) delete(data int) error {
	// Edit here
}

func (b *bst) isFull() bool {
	// Edit here
}

func (b *bst) isPerfect() bool {
	// Edit here
}

func (b *bst) isComplete() bool {
	// Edit here
}

func (b *bst) isBalanced() bool {
	// Edit here
}	
	`,
}
