package sorting

import "errors"

// MaxHeap is a defined struct to hold the list of tuples passed to it and associate required methods to sort it
// based on a specified column.
// The sorted flag specifies whether or not the result is ready to be returned.
type MaxHeap struct {
	list   []Tuple
	column int
	sorted bool
}

// NewHeap is a function that instantiates a new pointer to a heap with a list of tuples passed to it and a column index
// to sort these tuples according to it, and a sorted flag to specify whether or not the list is sorted and ready to be returned.
func NewHeap(list []Tuple, column int) *MaxHeap {
	var returnedHeap *MaxHeap = &MaxHeap{
		list:   list,
		column: column,
		sorted: false,
	}
	return returnedHeap
}

// heapify is an associated method to the MaxHeap struct, sorts out the binary tree of the tuples which are the elements
// of the heap, by swapping the root of the subtree with its left or right child if its value was smaller than them,
// which creates a sorted max-heap of tuples according to the column of comparison.
func (heap *MaxHeap) heapify(heapSize, subtreeRootIndex int) {
	var largest int = subtreeRootIndex
	var leftChild int = 2*subtreeRootIndex + 1
	var rightChild int = 2*subtreeRootIndex + 2

	if leftChild < heapSize && heap.list[largest][heap.column] < heap.list[leftChild][heap.column] {
		largest = leftChild
	}
	if rightChild < heapSize && heap.list[largest][heap.column] < heap.list[rightChild][heap.column] {
		largest = rightChild
	}
	if largest != subtreeRootIndex {
		heap.list[subtreeRootIndex], heap.list[largest] = heap.list[largest], heap.list[subtreeRootIndex]
		heap.heapify(heapSize, largest)
	}
}

// heapSort is an associated method with the struct MaxHeap, it feeds the method heapify with the correct heap size
// to sort upto and based on it, with each max-heap being structured, the root element (the largest value) is swapped
// with the last element (the smallest value), then reducing the size of the heap by one, until all elements are sorted
// in an increasing order.
func (heap *MaxHeap) heapSort() {
	var heapSize int = len(heap.list)
	for subtreeRootIndex := (heapSize / 2) - 1; subtreeRootIndex >= 0; subtreeRootIndex-- {
		heap.heapify(heapSize, subtreeRootIndex)
	}
	for newHeapSize := heapSize - 1; newHeapSize > 0; newHeapSize-- {
		heap.list[newHeapSize], heap.list[0] = heap.list[0], heap.list[newHeapSize]
		heap.heapify(newHeapSize, 0)
	}
	heap.sorted = true
}

func (heap *MaxHeap) getSortedList() ([]Tuple, error) {
	if heap.sorted {
		return heap.list, nil
	} else {
		var err error = errors.New("the passed list is not sorted yet")
		return nil, err
	}
}
