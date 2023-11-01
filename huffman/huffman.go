package huffman

type node struct {
	symbol rune
	frequency int
	left, right *node
}

type minHeap struct {
	itemCount int
	values []node
}

func initMinHeap() minHeap {
	return minHeap{}
}

func minHeapify(arr []node, size, smallestIndex int) {
	smallest := smallestIndex
	left := 2 * smallestIndex + 1
	right := 2 * smallestIndex + 2

	if left < size && arr[smallestIndex].frequency > arr[left].frequency {
		smallest = left
	}

	if right < size && arr[smallest].frequency > arr[right].frequency {
		smallest = right
	}

	if smallest != smallestIndex {
		arr[smallestIndex], arr[smallest] = arr[smallest], arr[smallestIndex]
		minHeapify(arr, size, smallest)
	}
	
}

func (h *minHeap) insert(frequency int, symbol rune, left, right  *node) {
	newNode := node{
		frequency: frequency,
		symbol: symbol,
		left: left,
		right: right,
	}
	h.values = append(h.values, newNode)
	h.itemCount++
	
	for i := (h.itemCount / 2) - 1; i > -1; i-- {
		minHeapify(h.values, h.itemCount, i)
	}
}

func (h *minHeap) delete() node {
	if h.itemCount == 0 {
		return node{}
	}
	val := h.values[0]
	h.itemCount--
	h.values[0] = h.values[h.itemCount]
	h.values = h.values[:h.itemCount]
	for i := (h.itemCount / 2) - 1; i > -1; i-- {
		minHeapify(h.values, h.itemCount, i)
	}
	return val
}

func tranverse(node *node, code string, huffmanCode map[string]string) {
	if node.left != nil {
		newCode := code + "0"
		tranverse(node.left, newCode, huffmanCode)
	}
	if node.right != nil {
		newCode := code + "1"
		tranverse(node.right, newCode, huffmanCode)
	}
	if node.right == nil && node.left == nil {
		huffmanCode[string(node.symbol)] = code
	}
}


func Huffman(input string) map[string]string {
	// Obtain the frequency of each character in the string
	characterFrequency := make(map[rune]int)
	for _, v := range input {
		characterFrequency[v] = characterFrequency[v] + 1
	}

	priorityQueue := initMinHeap()
	// Store the characters in an increasing order of frequency
	for symbol, frequency := range characterFrequency {
		priorityQueue.insert(frequency, symbol, nil, nil)
	}

	// Only one item should remain in the priority queue
	for priorityQueue.itemCount > 1 {
		leftChild := priorityQueue.delete()
		rightChild := priorityQueue.delete()
		// Add a new node whose
		// leftChild is the minimum frequency
		// rightChild is the second minimum frequency
		// data is the sum of the frequency of the left and right child
		priorityQueue.insert(
			leftChild.frequency + rightChild.frequency,
			0,
			&leftChild,
			&rightChild,
		)
	}

	huffmanCode := make(map[string]string)
	rootNode := priorityQueue.delete()
	// Tranverse the remaining node in the priority queue
	// In such a way the left edge gets 0 while the right one gets 1
	tranverse(&rootNode, "", huffmanCode)
	return huffmanCode
}