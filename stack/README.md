# Stack Explanation

Visualize a stack as a pile of plates on top of each other in such a way that the last plate to be inserted is the first plate to be retrieved (LIFO).

# Stack Implementation

There are five funcions publicly exposed:

1. Init(size)
> Creates a stack based on a particular size supplied by a user.
> The data type of the items on the stack is not restricted.
> It returns the newly created stack and an error (used to check if initialization is successful).

2. Push(item)
> It adds the item supplied by the user to the stack as along as the stack is not full.
> It returns an error (used to check if pushing is successful).
> It is a method of the initialized stack.

3. Pop()
> It removes the last inserted item from the stack as along as the stack is not empty.
> It returns the item or nil.
> It is a method of the initialized stack.

4. Peek()
> It returns the last inserted item from the stack as along as it is not empty (it does not remove it from the stack).
> It returns the item or nil.
> It is a method of the initialized stack.

5. IsEmpty()
> It specifies whether the stack contains items or not.
> It returns a boolean value.
> It is a method of the initialized stack.

6. IsFull()
> It specifies whether the stack is full based on the size used during initialization or not.
> It returns a boolean value.
> It is a method of the initialized stack.

# Example
```go

import "github.com/o-richard/DSA/stack"

                    .
                    .
                    .

myStack, err := stack.Init(2)
if err != nil {
    // There is a problem during initialization
}

err = myStack.Push(6)
if err != nil {
    // There is a problem during pushing
}

val := myStack.Peek() // Retuns 6

val = myStack.pop() // Retuns 6

if myStack.IsFull() {
    // The stack is full
}
if myStack.IsEmpty() {
    // The stack is empty
}

```

# Applications

1. Reverse a word.
2. Function call management in programming  languages.
3. Undo functionality.
4. Backtracking algorithms such as Depth First Search.
5. Expression evaluation; converting the value of expressions like 2 + 4 / 5 * (7 - 9) to prefix or postfix form from infix form.
6. Browser history enabling users to navigate back through their browsing history.
7. Memory management keeping track of execution of programs and the allocation of local variables.
8. Task scheduling in operating Systems; it helps in execution of processes and managing their state information.
9. Recursion

# Tests
