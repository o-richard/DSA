package prim

import "slices"

type Edge struct {
	Source, Destination string
	Weight int
}

type node struct {
	source, destination string
	data int
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

	if left < size && arr[smallestIndex].data > arr[left].data {
		smallest = left
	}

	if right < size && arr[smallest].data > arr[right].data {
		smallest = right
	}

	if smallest != smallestIndex {
		arr[smallestIndex], arr[smallest] = arr[smallest], arr[smallestIndex]
		minHeapify(arr, size, smallest)
	}
	
}

func (h *minHeap) insert(data int, source, destination string) {
	h.values = append(h.values, node{data: data, source: source, destination: destination})
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

func (h *minHeap) isEmpty() bool {
	return h.itemCount == 0
}

func Prim(graph map[string]map[string]int) ([]Edge, int) {
	var uniqueVertices []string
	for key, values := range graph {
		if !slices.Contains(uniqueVertices, key) {
			uniqueVertices = append(uniqueVertices, key)
		}
		for value := range values {
			if !slices.Contains(uniqueVertices, value) {
				uniqueVertices = append(uniqueVertices, value)
			}	
		}
	}

	var minCost int
	var result []Edge
	
	if len(uniqueVertices) == 0 {
		return result, minCost
	}

	visited := make(map[string]bool)
	priorityQueue := initMinHeap()
	priorityQueue.insert(0, uniqueVertices[0], uniqueVertices[0])

	for !priorityQueue.isEmpty() {
		minNode := priorityQueue.delete()

		if visited[minNode.destination] {
			continue
		}

		minCost += minNode.data
		visited[minNode.destination] = true

		newEdge := Edge{Source: minNode.source, Destination: minNode.destination, Weight: minNode.data}
		result = append(result, newEdge)

		for k, v := range graph[minNode.destination] {
			if !visited[k] {
				priorityQueue.insert(v, minNode.destination, k)
			}
		}
	}

	return result[1:], minCost
}