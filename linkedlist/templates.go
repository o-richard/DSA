package linkedlist

var Templates = map[string]string{
	"singleLinkedList": `package linkedlist

type sLL struct {
	// Edit here
}

func InitSLL() sLL {
	// Edit here
}

func (s *sLL) tranverse(_ bool) []int {
	// Edit here
}

func (s *sLL) search(searchData int) (int, bool) {
	// Edit here
}

func (s *sLL) insertion(inputData int, insertionIndex int) error {
	// Edit here
}

func (s *sLL) retrieval(retrievalIndex int, isDelete bool) (int, error) {
	// Edit here
}	
	`,
	"doubleLinkedList": `package linkedlist

type dLL struct {
	// Edit here
}

func InitDLL() dLL {
	// Edit here
}

func (d *dLL) tranverse(_ bool) []int {
	// Edit here
}

func (d *dLL) search(searchData int) (int, bool) {
	// Edit here
}

func (d *dLL) insertion(inputData int, insertionIndex int) error {
	// Edit here
}

func (d *dLL) retrieval(retrievalIndex int, isDelete bool) (int, error) {
	// Edit here
}
	`,
	"circularSingleLinkedList": `package linkedlist

type cSLL struct {
	// Edit here
}

func InitCSLL() cSLL {
	// Edit here
}

func (c *cSLL) tranverse(_ bool) []int {
	// Edit here
}

func (c *cSLL) search(searchData int) (int, bool) {
	// Edit here
}

func (c *cSLL) insertion(inputData int, insertionIndex int) error {
	// Edit here
}

func (c *cSLL) retrieval(retrievalIndex int, isDelete bool) (int, error) {
	// Edit here
}
	`,
	"circularDoubleLinkedList": `package linkedlist

type cDLL struct {
	// Edit here
}

func InitCDLL() cDLL {
	// Edit here
}

func (c *cDLL) tranverse(goLeftToRight bool) []int {
	// Edit here
}

func (c *cDLL) search(searchData int) (int, bool) {
	// Edit here
}

func (c *cDLL) insertion(inputData int, insertionIndex int) error {
	// Edit here
}

func (c *cDLL) retrieval(retrievalIndex int, isDelete bool) (int, error) {
	// Edit here
}
	`,
}
