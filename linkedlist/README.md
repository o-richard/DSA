# Linked List Explanation

### Single Linked List
Visualize a chain of carriages of a train where each carriage (node) has cargo/people (data) and a connection to the next carriage via a door (a pointer). Starting from the engine room (head) you can travel through the train by opening each door (following the pointers). The door on the last carriage will be open ended (end of the train).
**You can only move in a single direction from the engine room till the last carriage**

### Double Linked List
Relying on the visualization from the single linked list, imagine the carriages now have two doors in such a way you can travel to the last visited carriage (previous pointer) and also to the next carriage (next pointer).**You can move in a two directions from the engine room till the last carriage and back**

### Circular Single Linked List
Visualize a circular track with a series of runners in such a way that a runner (node) is holding a baton (pointer) and is passing it to the next runner (node). You can start at any runner (node) indefinitely and keep passing the baton since there is no 'starting' point or 'ending' point. **It is just a loop/continuous cycle with a traversal in a single direction**

### Circular Double Linked List
Visulaize the circular track from the circular single linked list, however, the runner (node) has two batons (pointers) that are passed to the runner in front of him (next pointer) and behind him (previous pointer) **It is just a loop/continuous cycle with a traversal on both directions**


# Linked List Implementation
### Single Linked List
There is one function currently exposed:

1. InitSLL()
> It creates a single linked list.
> The data type is restricted to integers.

### Double Linked List
There is one function currently exposed:

1. InitDLL()
> It creates a double linked list.
> The data type is restricted to integers.

### Circular Single Linked List
There is one function currently exposed:

1. InitCSLL()
> It creates a circular single linked list.
> The data type is restricted to integers.

### Circular Double Linked List
There is one function currently exposed:

1. InitCDLL()
> It creates a circular double linked list.
> The data type is restricted to integers.

### Shared (Interface for interaction)
There are four functions currently exposed:

1. Tranverse(linkedlist, direction)
> It tranverses the linked list in the specified direction.
> Only a circular double linked list allows for forward and backward traversals. The default option is forward for other linked lists.
> It returns an array of integers based on the diretion of traversal.
> The function is implemented as a method for the specified linked list.

2. Search(linkedlist, searchData)
> It searches for the search data in the specified linked list.
> It returns the index of the item in the linked list and whether the item was found or not.
> The function is implemented as a method for the specified linked list.

3. Insertion(linkedlist, input data, insertion index)
> It inserts a node containing the input data at the specified index.
> It supports both negative and positive insertion indices.
> It returns an error whether the insertion was successful or not.
> The function is implemented as a method for the specified linked list.

4. Retrieval(linkedlist, retrieval index, should delete?)
> It retrieves the node at the specified index.
> It supports both negative and positive insertion indices.
> It returns the data at the retrieval index and an error whether the retrieval was successful or not.
> The node is only deleted if the user species they would want to delete the node.
> The function is implemented as a method for the specified linked list.

# Example

### Single Linked List
```go
import "github.com/o-richard/DSA/linkedlist"
                        .
                        .
                        .
myList := linkedlist.InitSLL()
err := linkedlist.Insertion(&myList, 8, 0)
if err != nil {
    // There was an error during insertion.
}
val, found := linkedlist.Search(&myList, 7)
if found {
    // The value is found.
}
nums := linkedlist.Tranverse(&myList, true)
num, err := linkedlist.Retrieval(&myList, 0, false)
if err != nil {
    // There was an error during retrieval.
}
```
### Double Linked List
```go
import "github.com/o-richard/DSA/linkedlist"
                        .
                        .
                        .
myList := linkedlist.InitDLL()
err := linkedlist.Insertion(&myList, 8, 0)
if err != nil {
    // There was an error during insertion.
}
val, found := linkedlist.Search(&myList, 7)
if found {
    // The value is found.
}
nums := linkedlist.Tranverse(&myList, true)
num, err := linkedlist.Retrieval(&myList, 0, false)
if err != nil {
    // There was an error during retrieval.
}
```

### Circular Single Linked List
```go
import "github.com/o-richard/DSA/linkedlist"
                        .
                        .
                        .
myList := linkedlist.InitCSLL()
err := linkedlist.Insertion(&myList, 8, 0)
if err != nil {
    // There was an error during insertion.
}
val, found := linkedlist.Search(&myList, 7)
if found {
    // The value is found.
}
nums := linkedlist.Tranverse(&myList, true)
num, err := linkedlist.Retrieval(&myList, 0, false)
if err != nil {
    // There was an error during retrieval.
}
```

### Circular Double Linked List
```go
import "github.com/o-richard/DSA/linkedlist"
                        .
                        .
                        .
myList := linkedlist.InitCDLL()
err := linkedlist.Insertion(&myList, 8, 0)
if err != nil {
    // There was an error during insertion.
}
val, found := linkedlist.Search(&myList, 7)
if found {
    // The value is found.
}
nums := linkedlist.Tranverse(&myList, true)
num, err := linkedlist.Retrieval(&myList, 0, false)
if err != nil {
    // There was an error during retrieval.
}
```

# Big O Notation
| Operation                | Single Linked List | Double Linked List | Circular Single Linked List | Circular Double Linked List |
|--------------------------|--------------------|--------------------|-----------------------------|-----------------------------|
| Traversal                | O(n)               | O(n)               | O(n)                        | O(n)                        |
| Searching                | O(n)               | O(n)               | O(n)                        | O(n)                        |
| Insertion (at beginning) | O(1)               | O(1)               | O(1)                        | O(1)                        |
| Insertion (at end)       | O(n)               | O(1)               | O(1)                        | O(1)                        |
| Insertion (at middle)    | O(n)               | O(n)               | O(n)                        | O(n)                        |
| Deletion (at beginning)  | O(1)               | O(1)               | O(1)                        | O(1)                        |
| Deletion (at end)        | O(n)               | O(1)               | O(1)                        | O(1)                        |
| Deletion (at middle)     | O(n)               | O(n)               | O(n)                        | O(n)                        |



# Applications
1. Dynamic memory allocation
2. Undo/Redo functionality
3. Task management. It can maintain a list of tasks or to do items.
4. Forward and backward navigation in the browsing history.
5. Implemented in stack and queue.

# Tests
