package heap

import (
	"errors"
	"math"
)

// fHeap refers to a fibonacci heap
// fHeapNode refers to a fibonacci heap node
// degree refers to the number of child nodes for an efficient fibonacci heap

type fHeapNode struct {
	data, degree                               int
	isMarked, isVisited                        bool
	leftNode, rightNode, parentNode, childNode *fHeapNode
}

type fHeap struct {
	itemCount int
	isMaxHeap bool
	rootNode  *fHeapNode
}

func InitFHeap(isMaxHeap bool) fHeap {
	return fHeap{isMaxHeap: isMaxHeap}
}

func (f *fHeap) InsertFHeap(inputData int) {
	newNode := fHeapNode{data: inputData}
	if f.rootNode == nil {
		newNode.leftNode = &newNode
		newNode.rightNode = &newNode
		f.rootNode = &newNode
	} else {
		// Insertion to the left side of the root node
		f.rootNode.leftNode.rightNode = &newNode
		newNode.rightNode = f.rootNode
		newNode.leftNode = f.rootNode.leftNode
		f.rootNode.leftNode = &newNode

		// Update the root node if necessary
		if f.isMaxHeap {
			if newNode.data > f.rootNode.data {
				f.rootNode = &newNode
			}
		} else {
			if newNode.data < f.rootNode.data {
				f.rootNode = &newNode
			}
		}
	}
	f.itemCount++
}

func (f *fHeap) FindRootFHeap() (int, error) {
	if f.rootNode == nil {
		return 0, errors.New("there heap is empty")
	}
	return f.rootNode.data, nil
}

func (f *fHeap) UnionFHeap(newFHeap *fHeap) error {
	if newFHeap.rootNode == nil {
		return errors.New("the new heap should not be empty")
	}
	if f.rootNode == newFHeap.rootNode {
		return errors.New("the two heaps should be on different blocks of memory")
	}

	// Concatenating the root lists by inserting to the right of the main root list
	newFHeap.rootNode.rightNode.leftNode = f.rootNode.leftNode
	f.rootNode.leftNode.rightNode = newFHeap.rootNode.rightNode
	newFHeap.rootNode.rightNode = f.rootNode
	f.rootNode.leftNode = newFHeap.rootNode

	// Update the root node of the main heap if necessary
	if f.isMaxHeap {
		if newFHeap.rootNode.data > f.rootNode.data {
			f.rootNode = newFHeap.rootNode
		}
	} else {
		if newFHeap.rootNode.data < f.rootNode.data {
			f.rootNode = newFHeap.rootNode
		}
	}

	f.itemCount += newFHeap.itemCount
	return nil
}

func addChildToParent(child, parent *fHeapNode) {
	// Remove child from root list
	child.rightNode.leftNode = child.leftNode
	child.leftNode.rightNode = child.rightNode
	child.leftNode, child.rightNode = child, child
	child.parentNode = parent

	// Merge child to childlist of parent
	if parent.childNode == nil {
		parent.childNode = child
	} else {
		child.rightNode = parent.childNode
		child.leftNode = parent.childNode.leftNode
		parent.childNode.leftNode.rightNode = child
		parent.childNode.leftNode = child
	}
	parent.degree++
}

func consolidate(f *fHeap) {
	// Alternative: degree := (log(f.itemCount) / log(2)) + 1
	// NOTE: degree refers to the size of the array. The degree of the heap should be array size - 1
	_, degree := math.Frexp(float64(f.itemCount))

	arr := make([]*fHeapNode, degree)
	tracker := f.rootNode.leftNode
	for x := f.rootNode; ; {
		d := x.degree
		for arr[d] != nil {
			y := arr[d]

			if f.isMaxHeap {
				if x.data < y.data {
					x, y = y, x
				}
			} else {
				if x.data > y.data {
					x, y = y, x
				}
			}

			if y == tracker {
				tracker = x
			}
			addChildToParent(y, x)

			arr[d] = nil
			d++
		}
		arr[d] = x

		if x == tracker {
			break
		}

		x = x.rightNode
	}

	f.rootNode = nil
	// Come up with a new root node
	for _, tmp := range arr {
		if tmp != nil {
			if f.rootNode == nil {
				f.rootNode = tmp
			}
			if f.isMaxHeap {
				if tmp.data > f.rootNode.data {
					f.rootNode = tmp
				}
			} else {
				if tmp.data < f.rootNode.data {
					f.rootNode = tmp
				}
			}
		}
	}
}

func (f *fHeap) ExtractRootFHeap() (int, error) {
	if f.rootNode == nil {
		return 0, errors.New("the heap is empty")
	}

	// Add all child nodes to the root list of the heap
	if f.rootNode.childNode != nil {
		childNode := f.rootNode.childNode
		for {
			tmp := childNode.rightNode

			f.rootNode.leftNode.rightNode = childNode
			childNode.rightNode = f.rootNode
			childNode.leftNode = f.rootNode.leftNode
			f.rootNode.leftNode = childNode

			childNode.parentNode = nil
			childNode = tmp

			if tmp == f.rootNode.childNode {
				break
			}
		}
	}

	// Remove the root node
	f.rootNode.leftNode.rightNode = f.rootNode.rightNode
	f.rootNode.rightNode.leftNode = f.rootNode.leftNode

	val := f.rootNode.data

	if f.rootNode == f.rootNode.rightNode {
		// Remove the root node in case of one item
		f.rootNode = nil
	} else {
		f.rootNode = f.rootNode.rightNode
		consolidate(f)
	}

	f.itemCount--

	return val, nil
}

func findNode(searchPin *fHeapNode, searchData int) *fHeapNode {
	searchPin.isVisited = true

	if searchPin.data == searchData {
		searchPin.isVisited = false
		return searchPin
	}

	var node *fHeapNode
	if searchPin.childNode != nil {
		node = findNode(searchPin.childNode, searchData)
	}
	if !searchPin.rightNode.isVisited && node == nil {
		node = findNode(searchPin.rightNode, searchData)
	}

	// Ensure the SearchPin that called the function has reset its value for future findings.
	searchPin.isVisited = false
	return node
}

func cut(f *fHeap, child, parent *fHeapNode) {
	if child == child.rightNode {
		parent.childNode = nil
	} else {
		if child == parent.childNode {
			parent.childNode = child.rightNode
		}
	}

	child.leftNode.rightNode = child.rightNode
	child.rightNode.leftNode = child.leftNode

	parent.degree--

	child.rightNode = f.rootNode
	child.leftNode = f.rootNode.leftNode
	f.rootNode.leftNode.rightNode = child
	f.rootNode.leftNode = child

	if child.isMarked {
		child.isMarked = false
	}
	child.parentNode = nil
}

func cascadingCut(f *fHeap, parent *fHeapNode) {
	if parent.parentNode != nil {
		if !parent.isMarked {
			parent.isMarked = true
		} else {
			cut(f, parent, parent.parentNode)
			cascadingCut(f, parent.parentNode)
		}
	}
}

func (f *fHeap) ChangeKey(currentData, newData int) error {
	if f.rootNode == nil {
		return errors.New("the heap is empty")
	}
	exactNode := findNode(f.rootNode, currentData)
	if exactNode == nil {
		return errors.New("the node does not exist")
	}
	exactNode.data = newData

	// Parent node will be updated during cutting
	parent := exactNode.parentNode

	if f.isMaxHeap {
		if parent != nil && parent.data < newData {
			cut(f, exactNode, parent)
			cascadingCut(f, parent)
		}

		if newData > f.rootNode.data {
			f.rootNode = exactNode
		}
	} else {
		if parent != nil && parent.data > newData {
			cut(f, exactNode, parent)
			cascadingCut(f, parent)
		}
		if newData < f.rootNode.data {
			f.rootNode = exactNode
		}
	}
	return nil
}

func (f *fHeap) DeleteNode(currentData int) error {
	var newData int
	if f.isMaxHeap {
		newData = math.MaxInt64
	} else {
		newData = math.MinInt64
	}
	err := f.ChangeKey(currentData, newData)
	if err != nil {
		return err
	}
	f.ExtractRootFHeap()
	return nil
}
