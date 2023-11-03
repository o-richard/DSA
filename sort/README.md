# Sort Explanation and Implementation

### Bubble Sort

Imagine you have a row of unsorted cards. You start at one end and you repeatedly compare adjacent cards. Depending on the required order (ascending, descending), you swap the adjacents. The process is continued till no more swaps are needed.

### Selection Sort

Imagine a deck of cards spread out on a table. Depending on the required order (ascending, descending), you pick the smallest/largest card and put it in a separate pile. The process is repeated till all cards are sorted into the new pile.

### Insertion Sort

Imagine you have a row of unsorted cards. You pick one card at a time and insert it into its correct position in the sorted pile, shifting other cards if necessary.

### Merge Sort

Imagine you have a row of unsorted cards. You divide the cards into piles till each pile just contains a card. Then, you start merging the smaller piles into bigger sorted piles until you have a single sorted row of cards.

### Quicksort

Imagine you have a row of unsorted cards.  You randomly select a card to act as a pivot. You place all cards smaller than the pivot on one side of it while the cards larger than the pivot on the other sude. The process is repeated for the different decks until all cards are sorted.

### Counting Sort

Imagine you have a basket of apples of different colors. You create baskets for each color, count the apples of each color and then arrange them by color in the desired order.

### Radix Sort

Imagine you have a row of unsorted cards. You first sort them by their rightmost digits, creating piles of (0-9). Then, you sort the piles by the next digit to the left. The process is repated till all digits are considered.

### Heap Sort

Imagine you have a row of unsorted cards. You can create a min/max heap (depends on your configuration i.e ascending, descending). You remove the largest/smallest card number. The process is repeated till the heap becomes empty.

### Shell Sort

Imagine you have a row of unsorted cards, you oranize them into smaller piles. The smaller piles are sorted. The piles are mergedin such a way that the number of stacks are reduced. The process is repeated till one sorted stack remains.

### Bucket Sort

Imagine distributing balls of different sizes into buckets based on their range of sizes. Each bucket is sorted. The result of all the buckets is combined to form a sorted list.


# Complexities

| Sorting Algorithm    | Time Complexity (Best, Worst, Average) | Space Complexity | Stability |
|----------------------|-----------------------------------------|------------------|-----------|
| Bubble Sort          | O(n)        O(n^2)       O(n^2)          | O(1)             | Yes       |
| Selection Sort       | O(n^2)      O(n^2)       O(n^2)          | O(1)             | No        |
| Insertion Sort       | O(n)        O(n^2)       O(n^2)          | O(1)             | Yes       |
| Merge Sort           | O(n*log n)  O(n*log n)   O(n*log n)      | O(n)             | Yes       |
| Quicksort            | O(n*log n)  O(n^2)       O(n*log n)      | O(log n)         | No        |
| Counting Sort        | O(n+k)      O(n+k)       O(n+k)          | O(max)           | Yes       |
| Radix Sort           | O(n+k)      O(n+k)       O(n+k)          | O(max)           | Yes       |
| Heap Sort            | O(n*log n)  O(n*log n)   O(n*log n)      | O(1)             | No        |
| Shell Sort           | O(n*log n)  O(n^2)       O(n*log n)      | O(1)             | No        |


# Example

```go
import "github.com/o-richard/DSA/sort"
                    .
                    .
                    .
nums := []int {4, 3, 2, 1, 100},

fmt.Println(sort.BubbleSort(nums, true))
fmt.Println(sort.BubbleSort(nums, false))
fmt.Println(sort.CountingSort(nums))
fmt.Println(sort.HeapSort(nums, true))
fmt.Println(sort.HeapSort(nums, false))
fmt.Println(sort.InsertionSort(nums, true))
fmt.Println(sort.InsertionSort(nums, false))
fmt.Println(sort.MergeSort(nums, true))
fmt.Println(sort.MergeSort(nums, false))
fmt.Println(sort.QuickSort(nums, true))
fmt.Println(sort.QuickSort(nums, false))
fmt.Println(sort.RadixSort(nums))
fmt.Println(sort.SelectionSort(nums, true))
fmt.Println(sort.SelectionSort(nums, false))
fmt.Println(sort.ShellSort(nums, true))
fmt.Println(sort.ShellSort(nums, false))
```

# Applications

# Tests
