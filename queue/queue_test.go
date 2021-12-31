package queue

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	q := New[int](5, 1, 2, 3)
	q2 := New[int](5)

	exp := 1
	res := q.slice[0]
	if res != exp {
		t.Errorf("Expected queue first item to be: %v. Received: %v.", exp, res)
	}

	exp = 0
	res = q.head
	if res != exp {
		t.Errorf("Expected queue head to be: %v. Received: %v.", exp, res)
	}

	exp = 3
	res = q.tail
	if res != exp {
		t.Errorf("Expected queue tail to be: %v. Received: %v.", exp, res)
	}

	exp = 3
	res = q.size
	if res != exp {
		t.Errorf("Expected queue size to be: %v. Received: %v.", exp, res)
	}

	exp = 5
	res = q.maxSize
	if res != exp {
		t.Errorf("Expected queue maxSize to be: %v. Received: %v.", exp, res)
	}

	exp = 0
	res = q2.size
	if res != exp {
		t.Errorf("Expected queue size to be: %v. Received: %v.", exp, res)
	}

	exp = 5
	res = q2.maxSize
	if res != exp {
		t.Errorf("Expected queue maxSize to be: %v. Received: %v.", exp, res)
	}
}

func intSlicesEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestQueue_ToSlice(t *testing.T) {
	q := New[int](3, 1, 2, 3)
	q2 := New[int](3, 1, 2, 3)
	q2.Get()
	q2.Put(4) // TODO czy tak mozna zapetlac?
	q3 := New[int](0)

	exp := []int{1, 2, 3}
	res := q.ToSlice()
	if !intSlicesEqual(exp, res) {
		t.Errorf("Expected converted slice to be equal: %v. Received: %v.", exp, res)
	}

	exp = []int{2, 3, 4}
	res = q2.ToSlice()
	if !intSlicesEqual(exp, res) {
		t.Errorf("Expected converted slice to be equal: %v. Received: %v.", exp, res)
	}

	exp = []int{}
	res = q3.ToSlice()
	if !intSlicesEqual(exp, res) {
		t.Errorf("Expected converted slice to be equal: %v. Received: %v.", exp, res)
	}
}

func ExampleQueue_String() {
	q := New[int](0)
	fmt.Println(q)

	q = New[int](3, 1, 2, 3)
	fmt.Println(q)

	q = New[int](10, 1, 2, 3)
	fmt.Println(q)

	// Output:
	// queue[]
	// queue[1 2 3]
	// queue[1 2 3]
}

func TestQueue_Empty(t *testing.T) {
	q := New[int](0)
	if !q.Empty() {
		t.Error("Expected 0 maxSize queue to be empty.")
	}

	q = New[int](10)
	if !q.Empty() {
		t.Error("Expected 10 maxSize, no items queue to be empty.")
	}

	q = New[int](3, 1, 2, 3)
	if q.Empty() {
		t.Error("Expected queue with items to be non-empty.")
	}
}

func TestQueue_Full(t *testing.T) {
	q := New[int](3, 1, 2, 3)
	if !q.Full() {
		t.Error("Expected queue to be full.")
	}

	q = New[int](0)
	if q.Full() {
		t.Error("Expected queue not to be full.")
	}

	q = New[int](10, 1, 2, 3)
	if q.Full() {
		t.Error("Expected queue not to be full.")
	}
}

func TestQueue_Size(t *testing.T) {
	q := New[int](3, 1, 2, 3)
	exp := 3
	res := q.Size()
	if res != exp {
		t.Errorf("Expected queue size to be equal: %v. Received: %v.", exp, res)
	}

	q = New[int](0)
	exp = 0
	res = q.Size()
	if res != exp {
		t.Errorf("Expected queue size to be equal: %v. Received: %v.", exp, res)
	}

	q = New[int](10, 1, 2, 3)
	exp = 3
	res = q.Size()
	if res != exp {
		t.Errorf("Expected queue size to be equal: %v. Received: %v.", exp, res)
	}
}

func TestQueue_Get(t *testing.T) {
	q := New[int](3, 1, 2, 3)
	exp := 1
	res := q.Get()
	if res != exp {
		t.Errorf("Expected gotten value to be equal: %v. Received: %v.", exp, res)
	}

	exp = 2
	res = q.Get()
	if res != exp {
		t.Errorf("Expected gotten value to be equal: %v. Received: %v.", exp, res)
	}

	exp = 3
	res = q.Get()
	if res != exp {
		t.Errorf("Expected gotten value to be equal: %v. Received: %v.", exp, res)
	}
}

func TestQueue_Put(t *testing.T) {
	q := New[int](1)
	exp := 1
	q.Put(exp)
	res := q.Get()
	if res != exp {
		t.Errorf("Expected put value to be equal: %v. Received: %v.", exp, res)
	}

	exp = 5
	q.Put(exp)
	q.Put(7)
	res = q.Get()
	if res != exp {
		t.Errorf("Expected put value to be equal: %v. Received: %v.", exp, res)
	}

	exp = 7
	res = q.Get()
	if res != exp {
		t.Errorf("Expected put value to be equal: %v. Received: %v.", exp, res)
	}
}
