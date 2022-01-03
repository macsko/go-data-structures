package queue

import (
  "fmt"
  "testing"
)

func TestNew(t *testing.T) {
  q := New[int](5, 1, 2, 3)
  q2 := New[int](5)

  t.Run("Normal queue first item", testNewFunc(q.slice[0], 1, "first item"))
  t.Run("Normal queue head", testNewFunc(q.head, 0, "head"))
  t.Run("Normal queue tail", testNewFunc(q.tail, 3, "tail"))
  t.Run("Normal queue size", testNewFunc(q.size, 3, "size"))
  t.Run("Normal queue maxSize", testNewFunc(q.maxSize, 5, "maxSize"))
  t.Run("Empty queue size", testNewFunc(q2.size, 0, "size"))
  t.Run("Empty queue maxSize", testNewFunc(q2.maxSize, 5, "maxSize"))
}

func testNewFunc(res int, exp int, name string) func(*testing.T) {
  return func(t *testing.T) {
    if res != exp {
      t.Errorf("Expected queue %v to be equal: %v. Received: %v.", name, exp, res)
    }
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
  q2 := New[int](3, 1, 2, 3)
  q2.Get()
  q2.Put(4)

  t.Run("Normal queue", testQueue_ToSliceFunc(New[int](3, 1, 2, 3), []int{1, 2, 3}))
  t.Run("Extended queue", testQueue_ToSliceFunc(q2, []int{2, 3, 4}))
  t.Run("Empty queue", testQueue_ToSliceFunc(New[int](0), []int{}))
}

func testQueue_ToSliceFunc(q *Queue[int], exp []int) func(*testing.T) {
  return func(t *testing.T) {
    res := q.ToSlice()
    if !intSlicesEqual(exp, res) {
      t.Errorf("Expected converted slice to be equal: %v. Received: %v.", exp, res)
    }
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
  t.Run("Normal queue", testQueue_EmptyFunc(New[int](3, 1, 2, 3), false))
  t.Run("Empty queue", testQueue_EmptyFunc(New[int](0), true))
  t.Run("Empty queue with higher maxSize", testQueue_EmptyFunc(New[int](10), true))
}

func testQueue_EmptyFunc(q *Queue[int], exp bool) func(*testing.T) {
  return func(t *testing.T) {
    if q.Empty() != exp {
      if exp {
        t.Error("Expected queue to be empty.")
      } else {
        t.Error("Expected queue to be non-empty.")
      }
    }
  }
}

func TestQueue_Full(t *testing.T) {
  t.Run("Normal queue", testQueue_FullFunc(New[int](3, 1, 2, 3), true))
  t.Run("Empty queue", testQueue_FullFunc(New[int](0), false))
  t.Run("Queue with higher maxSize", testQueue_FullFunc(New[int](10, 1, 2, 3), false))
}

func testQueue_FullFunc(q *Queue[int], exp bool) func(*testing.T) {
  return func(t *testing.T) {
    if q.Full() != exp {
      if exp {
        t.Error("Expected queue to be full.")
      } else {
        t.Error("Expected queue to be non-full.")
      }
    }
  }
}

func TestQueue_Size(t *testing.T) {
  t.Run("Normal queue", testQueue_SizeFunc(New[int](3, 1, 2, 3), 3))
  t.Run("Empty queue", testQueue_SizeFunc(New[int](0), 0))
  t.Run("Queue with higher maxSize", testQueue_SizeFunc(New[int](10, 1, 2, 3), 3))
}

func testQueue_SizeFunc(q *Queue[int], exp int) func(*testing.T) {
  return func(t *testing.T) {
    res := q.Size()
    if res != exp {
      t.Errorf("Expected queue size to be equal: %v. Received: %v.", exp, res)
    }
  }
}

func TestQueue_Get(t *testing.T) {
  q := New[int](3, 1, 2, 3)

  t.Run("Normal queue 1", testQueue_GetFunc(q, 1))
  t.Run("Normal queue 2", testQueue_GetFunc(q, 2))
  t.Run("Normal queue 3", testQueue_GetFunc(q, 3))
}

func testQueue_GetFunc(q *Queue[int], exp int) func(*testing.T) {
  return func(t *testing.T) {
    res := q.Get()
    if res != exp {
      t.Errorf("Expected gotten value to be equal: %v. Received: %v.", exp, res)
    }
  }
}

func TestQueue_Put(t *testing.T) {
  q := New[int](3)
  q.Put(1)
  q.Put(3)
  q2 := New[int](3, 1, 2, 3)
  q2.Put(4)

  t.Run("Put 2 items - 1", testQueue_PutFunc(q, 1))
  t.Run("Put 2 items - 2", testQueue_PutFunc(q, 3))
  t.Run("Put 2 items into empty queue", testQueue_PutFunc(q2, 1))
}

func testQueue_PutFunc(q *Queue[int], exp int) func(*testing.T) {
  return func(t *testing.T) {
    res := q.Get()
    if res != exp {
      t.Errorf("Expected put value to be equal: %v. Received: %v.", exp, res)
    }
  }
}
