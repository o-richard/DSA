# Heap Explanation

For the explantion, implementation, example and applications of **a min and max heap**, refer to the [priority queue](../queue/README.md).


### Fibonacci Heap

**Based on your understanding of a min/max heap, apply it to understand a fibonacci heap. A major difference remains in a fibonacci heap, one cares about the order only after certain operations. It's not a necessity during some operations eg insertion.**

Imagine you have a deck of playing cards and you want to perform several operations on them such as inserting new cards, removing cards, changing cards, find the card with the highest value. 
> Visualize a fibonacci heap as an advanced card deck manger that optimizes these operations in a way that saves time and effort.
> You need to keep track on the minimum or maximum value based on the type of fibonacci heap. I'll refer to the item as a ***root node***

1. Insertion/Adding Cards
- Incase there are no cards, you can just place the card on the table and update the root node.
- During subsequent insertions, as you add a card, you can just place it on the table. The root node is only updated if the new item is lesser/more than the root node depending on your configuration.
- There is no need to organize the cards immediately, you can place them haphazardly.

2. Find the card with lowest/highest value.
- Incase you want to get the card with the lowest/highest value based on your configuration, you just refer to the root node.

3. Union/Combining Decks
- Incase you want to combine decks, you just have to connect their root nodes.
- The root node of the primary deck is updated only if the root node of the secondary deck is lesser/more compared to it (depends on the configuration of the heap).
- The order won't matter yet.

4. Removing the card with the highest/lowest value (root node)
- Remove the root node and reorganize the deck of cards to a certain order.

5. Changing the value of a card.
- You must locate the card you want to change, change its value and reorganie the deck of cards to maintain a certain order.

6. Removing a card
- You must locate the card, change the value of the card to the highest/lowest value possible depending on your configuration.
- The step above will ensure it will be the root node.
- Remove the root node.

# Heap Implementation

### Fibonacci Heap
There are 7 currently exposed methods:

1. InitFHeap(isMaxHeap?)
- Creates a fibonacci heap that satifies a min heap property or a max heap one.

2. InsertFHeap(inputData)
- Insert the data into the fibonacci heap.

3. FindRootFHeap()
- Obtains the value at the root node.
- It returns the value and an error (confirms the status of the operation)

4. UnionFHeap(newHeap)
- Combines two fibonacci heaps.
- It returns an error (confirms the status of the operation)

5. ExtractRootFHeap()
- Removes the value at the root node.
- It returns the values at the root and an error (confirms the status of the operation)

6. ChangeKey(current, new)
- Locates the node with the current value and updates its value to the new one.
- It returns an error (confirms the status of the operation)

7. DeleteNode(current)
- Deletes the node with the current value.
- It returns an error (confirms the status of the operation).

# Example

### Fibonacci Heap
```go
import "github.com/o-richard/DSA/heap"

                    .
                    .
                    .


myheap := heap.InitFHeap(false)
myheap1 := heap.InitFHeap(false)
myheap.InsertFHeap(6)
myheap.InsertFHeap(7)
myheap.InsertFHeap(8)
myheap1.InsertFHeap(1)
myheap1.InsertFHeap(2)
myheap1.InsertFHeap(3)
val, err := myheap.FindRootFHeap()
if err != nil {
    // There was an error.
}
err = myheap.UnionFHeap(&myheap1)
if err != nil {
    // There was an error
}
val, err = myheap.ExtractRootFHeap()
if err != nil {
    // There was an error.
}
err = myheap.ChangeKey(6, 9)
if err != nil {
    // There was an error
}
err = myheap.DeleteNode(6)
if err != nil {
    // There was an error
}

```
# Big O Notation

| Operation               | Fibonacci Heap |
|-------------------------|----------------|
| Insertion (Insert)      | O(1)           |
| Obtaining Minimum (FindMin) | O(1)       |
| Union (Merge)           | O(1)           |
| Deletion (Delete)       | O(log n)       |
| Decrease Key (DecreaseKey) | O(1)       |
| Delete Minimum (ExtractMin) | O(log n)   |

# Applications

### Fibonacci Heap
1. Dijkstra's algorithm
2. Optimization Problems

# Tests