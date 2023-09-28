package queue

var Templates = map[string]string{
	"circularQueue": `package queue

type circularQueue struct {
	// Edit here
}

func InitCircularQueue(size int) (circularQueue, error) {
	// Edit here
}

func (c *circularQueue) Enqueue(v interface{}) error {
	// Edit here
}

func (c *circularQueue) Dequeue() interface{} {
	// Edit here
}

func (c *circularQueue) Peek() interface{} {
	// Edit here
}

func (c *circularQueue) IsEmpty() bool {
	// Edit here
}

func (c *circularQueue) IsFull() bool {
	// Edit here
}
	`,
	"doubleEndedQueue": `package queue

type doubleEndedQueue struct {
	// Edit here
}

func InitDoubleEndedQueue(size int) (doubleEndedQueue, error) {
	// Edit here
}

func (d *doubleEndedQueue) AddAtFront(v interface{}) error {
	// Edit here
}

func (d *doubleEndedQueue) AddAtRear(v interface{}) error {
	// Edit here
}

func (d *doubleEndedQueue) DeleteAtFront() interface{} {
	// Edit here
}

func (d *doubleEndedQueue) DeleteAtRear() interface{} {
	// Edit here
}

func (d *doubleEndedQueue) IsEmpty() bool {
	// Edit here
}

func (d *doubleEndedQueue) IsFull() bool {
	// Edit here
}
	`,
	"priorityQueue": `package queue

type priorityQueue struct {
	// Edit here
}

func InitPriorityQueue(size int, isMaxHeap bool) (priorityQueue, error) {
	// Edit here
}

func (p *priorityQueue) Insert(input int) error {
	// Edit here
}

func (p *priorityQueue) Delete() (int, error) {
	// Edit here
}

func (p *priorityQueue) Peek() (int, error) {
	// Edit here
}

func (p *priorityQueue) IsEmpty() bool {
	// Edit here
}

func (p *priorityQueue) IsFull() bool {
	// Edit here
}	
	`,
}
