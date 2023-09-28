# Stack Explanation

Visualize a stack as a collection of plates on top of each other in such a way that the last plate to be inserted is the first plate to be retrieved (LIFO).

# Stack Implementation

There are five funcions publicly exposed:

1. Init(size)
> Creates a stack based on a particular size supplied by a user.
> The data type of the stack is not restricted.
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
> It is a method of the initialized stack.

6. IsFull()
> It specifies whether the stack is full based on the size used during initialization or not.
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

val := myStack.Peek()

val = myStack.pop()

if myStack.IsFull() {
    // The stack is full
}
if myStack.IsEmpty() {
    // The stack is empty
}

```

# Applications

1. To reverse a word.
2. To calculate the value of expressions like 2 + 4 / 5 * (7 - 9) by converting the expression to prefix or postfix form.
3. In browsers - The back button in a browser saves all the URLs you have visited previously in a stack.

# Tests