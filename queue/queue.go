package queue

type queue struct {
	front int
	rear int
	itemCount int
	maxSize int
	values []interface{}
}
