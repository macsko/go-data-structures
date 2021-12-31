// Package queue provides universal FIFO queue implementation using go1.18beta1 generics.
package queue

import "fmt"

// Queue is implemented as a circular array.
// It contains slice which represents an array with items, head representing index of first item of queue,
// tail which is an index of first free place in queue, size which contains a current number of items
// and maxSize which is a slice size.
// When size have to be higher than maxSize, the slice have to be extended and queue have to be reallocated.
// Construction with type interface{} allow storing items of multiple types in queue.
type Queue[T any] struct {
	slice   []T
	head    int
	tail    int
	size    int
	maxSize int
}

// New creates and returns a pointer to a new queue.
// It can be initially propagated with elems.
func New[T any](maxSize int, items ...T) *Queue[T] {
	if maxSize <= 0 {
		maxSize = 1
	}
	if maxSize <= len(items) {
		return &Queue[T]{items, 0, 0, len(items), len(items)}
	}
	slice := make([]T, maxSize)
	copy(slice, items)
	return &Queue[T]{slice, 0, len(items), len(items), maxSize}
}

// ToSlice returns slice converted from queue q.
func (q Queue[T]) ToSlice() []T {
	res := make([]T, q.size)
	if q.head < q.tail {
		copy(res, q.slice[q.head:q.tail])
	} else {
		copy(res, q.slice[q.head:q.maxSize])
		if q.tail != 0 {
			copy(res[q.maxSize-q.head:], q.slice[:q.tail])
		}
	}
	return res
}

// String returns string representation of queue q. Like:
// queue[1 2 3] represents queue with elements: 1, 2 and 3 (1 is a head, 3 is a tail)
// queue[] represents empty queue
func (q *Queue[T]) String() string {
	str := "queue"
	slice := q.ToSlice()
	str += fmt.Sprint(slice)
	return str
}

// Size returns actual number of items in the queue q.
func (q *Queue[T]) Size() int {
	return q.size
}

// Full reports whether queue q is full. It means that adding next item to queue will force resizing the queue.
func (q *Queue[T]) Full() bool {
	return q.size == q.maxSize
}

// Empty reports whether queue q is empty, so it does not contain any item.
func (q *Queue[T]) Empty() bool {
	return q.size == 0
}

// resize resizes the queue q, extending its maxSize twice.
// It should be called only when queue's slice is full.
func (q *Queue[T]) resize() {
	newSlice := make([]T, 2*q.maxSize)
	if q.head < q.tail {
		// If queue is not circular, newSlice is a copy of previous slice from head to tail,
		// but with extended length.
		copy(newSlice, q.slice[q.head:q.tail])
	} else if q.head == 0 {
		// If queue head is 0 and is not caught by previous if, the queue is full but not circular.
		// Then newSlice is a simple copy of previous slice, but with extended length.
		copy(newSlice, q.slice)
	} else {
		// If queue is circular, newSlice is a reallocated slice to start with head at index 0.
		// Firstly, it copies items from queue head to end of slice.
		copy(newSlice, q.slice[q.head:])
		// Secondly, it copies items from slice start to queue tail.
		copy(newSlice[q.size-q.head:], q.slice[:q.tail])
	}
	q.slice = newSlice
	q.head = 0
	q.tail = q.size
	q.maxSize *= 2
}

// Put puts a new item at the end of the queue q and if needed, resizes a queue.
func (q *Queue[T]) Put(item T) {
	if q.Full() {
		// Cannot insert new item to full queue. Have to resize it.
		q.resize()
	}
	q.slice[q.tail] = item
	q.tail = (q.tail + 1) % q.maxSize
	q.size++
}

// Get gets a head item from the queue q.
// It does not check whether queue is empty. It should be checked by a programmer calling Empty method first.
func (q *Queue[T]) Get() T {
	res := q.slice[q.head]
	q.head = (q.head + 1) % q.maxSize
	q.size--
	return res
}
