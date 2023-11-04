# Data Structures And Algorithms

## Prerequisites
1. A basic understanding of computers and basic programming concepts. 
2. Familiarization on the Go programming language.

## Usage
There is no structured learning path to use; once to wish to tackle a package - tackle it extensively.

```bash
# Build
go build

# Testing
# 1. A specific package
go test ./PACKAGE_NAME
# 2. The whole project
go test ./...

# Generate Templates
# 1. The whole project
<EXECUTABLE_FILE> -command=generate -directory=.
# 2. Specific directories
<EXECUTABLE_FILE> -command=generate -directory=bfs,dfs
# 3. A specific file
<EXECUTABLE_FILE> -command=generate -directory=bfs -file=bfs.go

# Restore Solutions
# 1. The whole project
<EXECUTABLE_FILE> -command=restore -directory=.
# 2. Specific directories
<EXECUTABLE_FILE> -command=restore -directory=bfs,dfs
# 3. A specific file
<EXECUTABLE_FILE> -command=restore -directory=bfs -file=bfs.go
```

### Recommedations
~~Develop a personal guide to use~~ Strictly stick to a personal guide you'll choose (Objetives vary, timelines to accomplish stuff vary, personalities differ).

~~Just rely on this repository~~ Advance your knowledge elsewhere and expand your viewpoints. There are platforms that provide challenges to tackle.


## Programming Language
### Why Go?
I find Go to be elegant (coming from the world of C, Python and Javascript).

### Recommendations
Implement data structures and algorithms in the C programming language if possible. (Gotta love the segfaults and dealing with raw memory with no leaks).

## Unimplemented Concepts

This repository is not extensive to cover everything on data structures and algorithms.

For example, some data structures are not implemented:

1. **Arrays** - Most languages have their own implementations of arrays. The underlying concept should be understood on how they work, their benefits, resizing concepts etc.
2. **Hashmaps** - Most langauages have their own implemenation of a hash map where a certain unique key maps to a particular value. The underlying concept should be understood on how they work, various hash functions, dealing with hash collisions, benefits etc. **Some concepts on hash functions are covered in [Rabin-Karp Algorithm](robinkarp/README.md)**

## Tests

Tests are implemented on every package to test the implementations of various data structures and algorithms.

Test cases are not extensive (Explanation on the README located at package level under the Test Section).

## Contributions

Contributions are welcome.

Suggestions:
- New packages should be added only if the existing implementation don't cover the concept.
- Modification of existing packages or additional of new packages should have accompanying working tests.

## Caution
**This work should not be taken as a library where you can import implementations of data structures and algorithms.**

This work should be a learning path for one to practise implementations of various data structures and algorithms so that they can use them in their own applications.

## Credits
- [https://www.scaler.com/topics/data-structures/](https://www.scaler.com/topics/data-structures/)
- [https://www.programiz.com/dsa](https://www.programiz.com/dsa)
- [https://www.cs.usfca.edu/~galles/visualization/Algorithms.html](https://www.cs.usfca.edu/~galles/visualization/Algorithms.html)
- [https://www.geeksforgeeks.org/learn-data-structures-and-algorithms-dsa-tutorial](https://www.geeksforgeeks.org/learn-data-structures-and-algorithms-dsa-tutorial)
- [https://www.geeksforgeeks.org/fundamentals-of-algorithms/](https://www.geeksforgeeks.org/fundamentals-of-algorithms/)