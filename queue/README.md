# Queue Expalantion

### Circular Queue

Visualize a circular queue as people queueing for a ride in a theme park in such a way that first come first serve (FIFO) and the person who just finished their ride can join the queue of people at the back of the queue.

### Priority Queue

Visualize a priority queue as patients arriving and getting served in a medical institution; the people who have critical conditions are given a higher priority and served first. FIFO is not necessary.

### Double Ended Queue

Visualize a double ended queue as a deck of cards stacked on each other in such a way that you can add or remove a card from the start or the end of the deck.

# Queue Implemenetation
### Circular Queue
There are six functions currently exposed:

1. InitCircularQueue(size)
> Creates a circular queue based on a particular size supplied by a user.
> The data type of the items of the circular queue is not restricted.
> It returns the newly created circular queue and an error (used to check if initialization is successful).

2. Enqueue(item)
> It adds an item supplied by the user to the circular queue as along as the queue is not full.
> It returns an error (used to check if enqueueing is successful).
> It is a method of the initialized circular queue.

3. Dequeue()
> It removes the first inserted item from the circular queue as along as the queue is not empty.
> It returns the item or nil.
> It is a method of the initialized circular queue.

4. Peek()
> It returns the first inserted item from the circular queue as along as the queue is not empty (it does not remove it).
> It returns the item or nil.
> It is a method of the initialized circular queue.

5. IsEmpty()
> It specifies whether the circular queue contains any item.
> It returns a boolean value (true or false)
> It is a method of the initialized circular queue.

6. IsFull()
> It specifies whether the circular queue is full based on the size specified by the user during initialization.
> It returns a boolean value (true or false)
> It is a method of the initialized circular queue.

### Priority Queue

There are 6 functions currentyly exposed

1. InitPriorityQueue(size, isMaxHeap)
> This creates a priority queue with the size specified and denotes if it should be a max heap (parent has the maximum value to its children) or a min heap (parent has the minimum value to its children).
> The data type of the priority queue is restricted to integers for easy comparison.
> It  returns the newly created priority queue and an error (used to check if initialization is successful).

2. Insert(item)
> This inserts an item (integer) to the priority queue as along as it is not full.
> The priority queue is heapified so that to maintain it being a max heap or a min heap depending on the specifications during initialization.
> This is a method on the created priority queue.
> It returns an error (used to check if insertion was successful).

3. Delete()
> This removes the root item of the while priority queue as along as the queue is not empty.
> The priority queue is heapified so that to maintain it being a max heap or a min heap depending on the specifications during initialization.
> This is a method on the created priority queue.
> It returns the item and an error (used to check if it is successful; queue was not empty).

4. Peek()
> This returns the root item of the while priority queue as along as the queue is not empty (does not remove it).
> This is a method on the created priority queue.
> It returns the item and an error (used to check if it is successful; queue was not empty).

5. IsEmpty()
> Specifies whether the priority queue is empty or not based on the present number of items.
> This is a method on the created priority queue.

6. IsFull()

> Specifies whether the priority queue is full or not based on the size specified during initialization.
> This is a method on the created priority queue.


### Double Ended Queue

There are 7 operations exposed

1. InitDoubleEndedQueue(size)
> Creates the double ended queue wuth the size specified by the user.
> The data type of the items of the double ended queue is not restrictive.
> It returns the double ended queue and an error (used to check if initialization is successful)

2. AddAtFront(item)
> It adds the item to the start of the queue as along as the queue is not full.
> This is a method on the created double ended queue.
> It returns an error (used to check if addition is successful).

3. AddAtRear(item)
> It adds the item to the end of the queue as along as the queue is not full.
> This is a method on the created double ended queue.
> It returns an error (used to check if addition is successful).

4. DeleteAtFront()
> It removes the item present at the start of the queue as along as the queue is not empty.
> It returns the item or nil (if empty queue).
> This is a method on the created double ended queue.

5. DeleteAtRear()
> It removes the item present at the end of the queue as along as the queue is not empty.
> It returns the item or nil (if empty queue).
> This is a method on the created double ended queue.

6. IsEmpty()
> Specifies whether the double ended queue is empty or not based on the current number of items.
> This is a method on the created double ended queue.

7. IsFull()
> Specifies whether the double ended queue is full or not based on the size specified during initialization.
> This is a method on the created double ended queue.

# Example
### Circular Queue
```go

import "github.com/o-richard/DSA/queue"

                    .
                    .
                    .

myQueue, err := InitCircularQueue(4)
if err != nil {
    // Something wrong happened during initialization
}

// Errors should be checked to ensure there was nothing wrong with enqueue
err = myQueue.Enqueue("a")
err = myQueue.Enqueue("b")
err = myQueue.Enqueue("c")
err = myQueue.Enqueue("d")

// Retutns "a"; the first item on the queue
val := myQueue.Dequeue()
// Retutns "a"; the first item on the queue
val = myQueue.Peek()

if myQueue.IsFull() {
    // The  queue is full
}

if myQueue.IsEmpty() {
    // The  queue is empty
}
```

### Priority Queue
```go

import "github.com/o-richard/DSA/queue"

                    .
                    .
                    .

myQueue, err := InitPriorityQueue(4, true)
if err != nil {
    // Something wrong happened during initialization
}

// Errors should be checked to ensure there was nothing wrong with insert
err = myQueue.Insert(1)
err = myQueue.Insert(2)
err = myQueue.Insert(3)
err = myQueue.Insert(4)

// Retutns 4; the root node on the max heap. Check on any error.
var val interface{}
val, err = myQueue.Delete()
// Returns 4; the root node on the max heap.  Check on any error.
val, err = myQueue.Peek()

if myQueue.IsFull() {
    // The  queue is full
}

if myQueue.IsEmpty() {
    // The  queue is empty
}
```

### Double Ended Queue
```go

import "github.com/o-richard/DSA/queue"

                    .
                    .
                    .

myQueue, err := InitDoubleEndedQueue(5)
if err != nil {
    // Something wrong happened during initialization
}

// Errors should be checked to ensure there was nothing wrong with insert
err = myQueue.AddAtFront(8)
err = myQueue.AddAtRear(5)
err = myQueue.AddAtFront(7)
err = myQueue.AddAtFront(10)
err = myQueue.AddAtRear(11)
err = myQueue.AddAtRear(7)
val := myQueue.DeleteAtRear()
val = myQueue.DeleteAtFront()

if myQueue.IsFull() {
    // The  queue is full
}

if myQueue.IsEmpty() {
    // The  queue is empty
}
```


# Applications
### Circular Queue
1. CPU scheduling. It can be used to implement round-robin scheduling algorithms for CPU time allocation.
2. Memory management. It can be used in managing a buffer of fixed size, like in streaming media applications to store and process a continuous stream of data.
3. Traffic management for example printer queue.

### Priority Queue
1. Dijkstra's algorithm. Nodes with the lowest cost are explored first.
2. Data compression in Huffman code. Characters with higher frequencies have higher priorities.
3. For implementing stack.
4. For load balancing and interrupt handling in an operating system.
5. Job scheduling.

### Double Ended Queue
1. In undo operations on software.
2. To store history in browsers.
3. Implementing both stacks and queues.
4. Palindrome checking. It can be used in checking whether a given string is a palindrome by using a deque to compare characters from both ends.

# Tests