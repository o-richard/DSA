package stack

var Templates = map[string]string{
	"stack": `package stack

type stack struct {
	// Edit here
}

func Init(size int) (stack, error) {
	// Edit here
}

func (s *stack) Push(v interface{}) error {
	// Edit here
}

func (s *stack) Pop() interface{} {
	// Edit here
}

func (s *stack) Peek() interface{} {
	// Edit here
}

func (s *stack) IsEmpty() bool {
	// Edit here
}

func (s *stack) IsFull() bool {
	// Edit here
}	
	`,
}
