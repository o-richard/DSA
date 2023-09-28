# Rabin-Karp Algorithm Explanation

Rabin-Karp Algorithm is used for searching/matching patterns in a text using a hash function. An efficient hash function will quickly eliminate positions where there can't be a match to speed up the search.

# Rabin-Karp Algorithm Implementation

Imagine you are looking for a specific word in a long piece of text i.e book and you want to find all occurences of the word.

You will need to come up with a unique hash values for your phrases, i.e the word car may be the value 5.

1. Calculate the hash value of the specific word and the hash value of the phrase at the start of the text (The phrase length does not exceed the pattern length)
2. Compare the hash values
    - Incase of a possible match, verify there is an actual match character by character. If there is an actual match, record the position.
3. Move the word to the right side of the starting phrase and recompute the hashes. 
4. Repeat step 2 and 3 until you have gone through the whole text.



# Example

```go
import "github.com/o-richard/DSA/robinkarp"
                        .
                        .
                        .
// Obtain the indices where matches are found
fmt.Println(robinkarp.Robinkarp("AABAABAAB", "AAB"))
```

# Applications

- Searching a string in a bigger text
- Pattern matching

# Tests