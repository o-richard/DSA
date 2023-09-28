# Search Explanation and Implementation

Imagine you are looking for a book in a library.

### Linear Search

**There is no information on how the books are arranged.** You will start to look for the book shelf by shelf until you find it or you exhaust all the shelves.

### Binary Search

**The books are arranged alphabetically according to the books title.** You will start at the middle of the library and based on the title of the book you are searching for, you will go to the left or right half of the library. You will repeat the process till you remain with a book which may or may not be the one you are looking for.

### Interpolation Search

**The books are not only arranged alphabetically by the title but also the are sorted by the publication date** Based on the title of your book and publication date, you can make an educated quess on which particular section of the library is relevant for your search. Your further fine tune your estimate and search on the relevant section. You will repeat the process till you find the book if it exists.

# Example

```go

import "github.com/o-richard/DSA/search"
                        .
                        .
                        .

fmt.Println(search.BinarySearch([]int {1, 2, 3, 4}, 5))
fmt.Println(search.LinearSearch([]int {1, 2, 3, 4}, 5))
```

# Applications

- Searching operations

# Tests
