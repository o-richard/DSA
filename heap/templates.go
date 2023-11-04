package heap

var Templates = map[string]string{
	"fibonacciHeap": `package heap

type fHeap struct {
	// Edit here
}

func InitFHeap(isMaxHeap bool) fHeap {
	// Edit here
}

func (f *fHeap) InsertFHeap(inputData int) {
	// Edit here
}

func (f *fHeap) FindRootFHeap() (int, error) {
	// Edit here
}

func (f *fHeap) UnionFHeap(newFHeap *fHeap) error {
	// Edit here
}

func (f *fHeap) ExtractRootFHeap() (int, error) {
	// Edit here
}

func (f *fHeap) ChangeKey(currentData, newData int) error {
	// Edit here
}

func (f *fHeap) DeleteNode(currentData int) error {
	// Edit here
}	
	`,
}
