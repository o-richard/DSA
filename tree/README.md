# Trees Explanation

### Binary Search Tree (BST)

Imagine you have a library of books and you would like to organize them in a way that allows for quick searching. You start with an empty shelf, each time you add a book, compare the new book's title with the current book in the shelf. If it is smaller, you go to the left and vice versa if it is larger. The steps are repeated till the book is inserted during an insertion operation or found during a retrieval operation.

### AVL Tree

Consider the library in the section **'Binary Search Tree'**, but this time, you want to ensure the shelf remains balanced. As you add/remove a book, a balance is ensured. In case of an imbalance, the shelf's arrangement is adjusted to maintain balance.

### Red-Black Tree

Consider the library in the section **'AVL Tree'**, but this time instead of perfectly balancing the shelves, a specific set of rules that ensure a reasonable level of balance are used. The books are assigned the colors red or black. The set of rles are applied when adding/removing a book,

### B Tree

COnsider a large bookshop where you need to organize books that won't fit in a single shelf. You can crete a system of multiple shelf where each shelf can hold a fixed number of books. Incase a shelf is full a new oneis created. The shelves are organized in a tree like hierarchical structure that can handle a massive number of books efficiently.

### B+ Tree

Extend thebookshope in the section **'B Tree'**, but this time, only put the titles on the shelves. The actual books are stored on a separate location. The titles are organized in a way that allows efficient searching.

### Merkle Tree

Imagine you want to ensure the integrity of the books in your library (or verify the books in your library). You create a system where you take every neighbouring books, combine their title and create a new title for the pair. The process is repeated till you remain with a single title representing the whole library (the library fingerprint). If someone changes even one book, the fingerprint of the library won't match.

# Trees Implementation

### Binary Search Tree (BST)
There is one function currently exposed:

1. InitBST()
> It creates a Binary Search Tree.
> The data type is restricted to integers.

### AVL Tree
There is one function currently exposed:

1. InitAVL()
> It creates an AVL tree.
> The data type is restricted to integers.

### Shared (Interface for interaction)
There are eight functions currently exposed:

1. Insert(tree, data)
> It inserts data into a tree
> The function is implemented as a method for the specified tree.

2. Search(tree, data)
> It searches the tree
> It returns whether the item is found or not
> The function is implemented as a method for the specified tree.

3. Tranverse(tree, tranverseType)
> It tranverses the tree
> TranverseType options are preorder, inorder, postorder
> It returns the order of the items based on the tranversal type
> The function is implemented as a method for the specified tree.

4. Delete(tree, data)
> It deletes the specified data from the tree
> It returns the error specifying if the deletion is successful
> The function is implemented as a method for the specified tree.

5. IsFull(tree)
> It specifies whether the tree is a full binary tree
> The function is implemented as a method for the specified tree.

6. IsPerfect(tree)
> It specifies whether the tree is a perfect binary tree
> The function is implemented as a method for the specified tree.

7. IsComplete(tree)
> It specifies whether the tree is a complete binary tree
> The function is implemented as a method for the specified tree.

8. IsBalanced(tree)
> It specifies whether the tree is a balanced binary tree
> The function is implemented as a method for the specified tree.

# Example

### Binary Search Tree (BST)
```go
import "github.com/o-richard/DSA/tree"
                .
                .
                .
mytree := tree.InitBST()
nums := []int{7,8,9,1,2,3,4,5,7,8,8,8,10,11, -100}	
for _, num := range nums {
    tree.Insert(&mytree, num)
}
for _, num := range nums {
    err := tree.Delete(&mytree, num)
    if err != nil {
        // An error occured
    }
}
for _, num := range nums {
    tree.Insert(&mytree, num)
}
for _, num := range nums {
    isFound := tree.Search(&mytree, num)
    if !isFound {
        // Absent item
    }
}
isBalanced := tree.IsBalanced(&mytree)
if !isBalanced {
    // Unbalanced tree
}
order, err:= tree.Tranverse(&mytree, "test")
if err != nil {
    // Unsupported type
}
```

### AVL Tree
```go
import "github.com/o-richard/DSA/tree"
                .
                .
                .
mytree := tree.InitAVL()
nums := []int{7,8,9,1,2,3,4,5,7,8,8,8,10,11, -100}	
for _, num := range nums {
    tree.Insert(&mytree, num)
}
for _, num := range nums {
    err := tree.Delete(&mytree, num)
    if err != nil {
        // An error occured
    }
}
for _, num := range nums {
    tree.Insert(&mytree, num)
}
for _, num := range nums {
    isFound := tree.Search(&mytree, num)
    if !isFound {
        // Absent item
    }
}
isBalanced := tree.IsBalanced(&mytree)
if !isBalanced {
    // Unbalanced tree
}
order, err:= tree.Tranverse(&mytree, "test")
if err != nil {
    // Unsupported type
}
```

# Applications

- Indexing in the database
- Dynamic sorting
- Searching
- Merkle Tree: Enuring data integrity and verification

# Tests

The functions IsComplete, IsFull and IsPerfect have not tested
