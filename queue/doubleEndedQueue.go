package queue

import "errors"

type doubleEndedQueue struct {
	front int
	rear int
	itemCount int
	maxSize int
	values []interface{}
}

func InitDoubleEndedQueue(size int) (doubleEndedQueue, error) {
	if size < 1 {
		return doubleEndedQueue{}, errors.New("the size should have a minimum of 1")
	}

	var d doubleEndedQueue
	d.maxSize = size
	d.rear = -1
	d.values = make([]interface{}, size)
	return d, nil
}

func (d *doubleEndedQueue) AddAtFront(v interface{}) error {
	if d.itemCount == d.maxSize {
		return errors.New("the queue is full")
	}

	d.itemCount++
	if d.front < 1 {
		d.front = d.maxSize - 1
	} else {
		d.front--
	}
	d.values[d.front] = v
	return nil
}

func (d *doubleEndedQueue) AddAtRear(v interface{}) error {
	if d.itemCount == d.maxSize {
		return errors.New("the queue is full")
	}
	d.itemCount++
	if d.rear == d.maxSize {
		d.rear = 0
	} else {
		d.rear++
	}
	d.values[d.rear] = v
	return nil
}

func (d *doubleEndedQueue) DeleteAtFront() interface{} {
	if d.itemCount == 0 {
		return nil
	}

	val := d.values[d.front]
	d.values[d.front] = nil

	if d.front == d.maxSize - 1 {
		d.front = 0
	} else {
		d.front++
	}
	
	d.itemCount--
	return val
}

func (d *doubleEndedQueue) DeleteAtRear() interface{} {
	if d.itemCount == 0 {
		return nil
	}
	
	val := d.values[d.rear]
	d.values[d.rear] = nil

	if d.rear == 0 {
		d.rear = d.maxSize - 1
	} else {
		d.rear--
	}

	d.itemCount--
	return val
}

func (d *doubleEndedQueue) IsEmpty() bool {
	return d.itemCount == 0
}

func (d *doubleEndedQueue) IsFull() bool {
	return d.itemCount == d.maxSize
}
