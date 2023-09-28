package queue

import "errors"

type circularQueue struct {
	front int
	rear int
	itemCount int
	maxSize int
	values []interface{}
}

func InitCircularQueue(size int) (circularQueue, error) {
	if size < 1 {
		return circularQueue{}, errors.New("the size should have a minimum of 1")
	}

	var c circularQueue
	c.maxSize = size
	c.values = make([]interface{}, size)
	return c, nil
}

func (c *circularQueue) Enqueue(v interface{}) error {
	if c.itemCount == c.maxSize {
		return errors.New("the queue is full")
	}
	
	c.values[c.rear] = v
	c.rear++
	if c.rear == c.maxSize {
		c.rear = 0
	}
	c.itemCount++
	return nil
}

func (c *circularQueue) Dequeue() interface{} {
	if c.itemCount == 0 {
		return nil
	}

	val := c.values[c.front]
	c.values[c.front] = nil
	c.front++
	if c.front == c.maxSize {
		c.front = 0
	}
	c.itemCount--
	return val
}

func (c *circularQueue) Peek() interface{} {
	if c.itemCount == 0 {
		return nil
	}
	return c.values[c.front]
}

func (c *circularQueue) IsEmpty() bool {
	return c.itemCount == 0
}

func (c *circularQueue) IsFull() bool {
	return c.itemCount == c.maxSize
}