package dijkstra

import (
	"math"
	"slices"
)

type node struct {
	name string
	data int
}

type minHeap struct {
	itemCount int
	values    []node
}

func initMinHeap() minHeap {
	return minHeap{}
}

func minHeapify(arr []node, size, smallestIndex int) {
	smallest := smallestIndex
	left := 2*smallestIndex + 1
	right := 2*smallestIndex + 2

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

func (h *minHeap) insert(data int, name string) {
	h.values = append(h.values, node{data: data, name: name})
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

func Dijkstra(graph map[string]map[string]int, source string) map[string]int {
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

	distance := make(map[string]int)
	for _, v := range uniqueVertices {
		distance[v] = math.MaxInt
	}
	distance[source] = 0
	priorityQueue := initMinHeap()
	priorityQueue.insert(0, source)

	for !priorityQueue.isEmpty() {
		minNode := priorityQueue.delete()

		for vertex, weight := range graph[minNode.name] {
			if distance[vertex] > minNode.data+weight {
				distance[vertex] = minNode.data + weight
				priorityQueue.insert(distance[vertex], vertex)
			}
		}
	}

	return distance
}
