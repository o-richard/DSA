# Huffman Coding Explanation

Huffman Coding is a method for lossless data compression which is meant to represent data in a more efficient way without losing any information. More common symbols have shorter codes.

# Huffman Coding Implementation

Imagine you have a list of characters.

Bulding a Huffman Tree
1. Count the frequecy of each character.
2. Make each unique character as a leaf node.
3. Create a new node that combines the two least frequent characters.
4. Insert the new node in the tree.
5. Repeat 3 - 5 for all unique characters.

Encoding characters
1. Starting from the root of the tree follow the path to the character you wish to encode. The path to the  left edge is a 0 while for the right one is a 1.

Decoding characters
1. Starting from the root of the tree and the encoded bits, the path will lead one to the character.

# Example

```go
"github.com/o-richard/DSA/huffman"
                    .
                    .
                    .
// Obtain the symbols present and their respective codes
fmt.Println(huffman.Huffman("BCAADDDCCACACAC"))
```

# Applications



# Tests