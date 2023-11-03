# Longest Common Subsequence Explanation

Longest Common Subsequence is the longest subsequence common to all the given sequences, provided that the elements of the subsequence are not required to occupy consecutive positions within the original sequences (However, the order of occurrence is still preseved).

Example One.
John and Doe have a set of books.
John - Book A, Book C, Book B
Doe - Book A, Book B.

The longest common subsequence is "Book A, Book B"

Example Two.
John and Doe have a set of books.
John - Book A, Book C, Book B
Doe - Book B, Book A

The longest common subsequence is "Book A" or "Book B"

Example Three.
John and Doe have a set of books.
John - Book A, Book C, Book B, Book D
Doe - Book B, Book A, Book D

The longest common subsequence is "Book A, Book D"


# Longest Common Subsequence Implementation

1. Based on the example of books, come up with a n by m table where n refers to the number of books in John's set + 1 while m refers to the number of books in Doe's set + 1.
2. Fill the first column and row with zeros.
3. Fill the whole table by the following logic
    - Incase the book corresponding to the current row and current column are matching, fill the current cell by adding one to the diagonal element. Point an arrow to the diagonal cell.
    - Incase there is no match, take the maximum value  from the previous column and previous row. Point an arrow to the chosen cell.
4. The value at the last column and last row will have the length of the longest common subsequence.
5. For the longest common subsequence of books, follow the arrows from the cell of the last row and column

# Example

```go
import "github.com/o-richard/DSA/longestcommonsubsequence"
                            .
                            .
                            .
longestcommonsubsequence.LCS("ACADB", "CBDA")
```

# Applications
- Compressing genome resequencing data
- Text comparisons

# Tests