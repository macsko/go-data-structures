package set

import (
  "fmt"
  "testing"
)

func TestNew(t *testing.T) {
  set := New[int](1, 2, 3)
  if len(*set) != 3 {
    t.Error("Expected 3 items in set. Received:", len(*set))
  }

  for e := range *set {
    if e != 1 && e != 2 && e != 3 {
      t.Error("Expected set to contain only 1, 2 and 3. Received:", e)
    }
  }
}

func slicesContentEqual[K comparable](a []K, b []K) bool {
  if len(a) != len(b) {
    return false
  }

outerLoop:
  for _, v := range a {
    for _, w := range b {
      if v == w {
        continue outerLoop
      }
    }
    return false
  }
  return true
}

func TestSet_ToSlice(t *testing.T) {
  t.Run("set[1, 2, 3]", testSet_ToSliceFunc(New[int](1, 2, 3), []int{1, 2, 3}))
  t.Run("set[3, 4]", testSet_ToSliceFunc(New[int](3, 4), []int{3, 4}))
}

func testSet_ToSliceFunc(set *Set[int], exp []int) func(*testing.T) {
  return func(t *testing.T) {
    res := set.ToSlice()
    if !slicesContentEqual(exp, res) {
      t.Errorf("Expected converted slice to be equal: %v. Received: %v.", exp, res)
    }
  }
}

func ExampleSet_String() {
  set := New[int]()
  fmt.Println(set)
  set = New[int](1)
  fmt.Println(set)

  // Output:
  // set[]
  // set[1]
}

func TestSet_Has(t *testing.T) {
  t.Run("Set contains", testSet_HasFunc(New[int](1, 2, 3), 2, true))
  t.Run("Set don't contain", testSet_HasFunc(New[int](1, 2, 3), 9, false))
}

func testSet_HasFunc(set *Set[int], num int, exp bool) func(*testing.T) {
  return func(t *testing.T) {
    res := set.Has(num)
    if res != exp {
      t.Errorf("Expected Has(%v) to return %v. Received: %v.", num, exp, res)
    }
  }
}

func TestSet_Equal(t *testing.T) {
  t.Run("Sets equal", testSet_EqualFunc(New[int](1, 2, 3), New[int](1, 2, 3), true))
  t.Run("Sets not equal", testSet_EqualFunc(New[int](1, 2, 3), New[int](1, 2), false))
}

func testSet_EqualFunc(set1 *Set[int], set2 *Set[int], exp bool) func(*testing.T) {
  return func(t *testing.T) {
    if set1.Equal(set2) != exp {
      if exp {
        t.Errorf("Expected sets %v and %v to be equal.", set1, set2)
      } else {
        t.Errorf("Expected sets %v and %v not to be equal.", set1, set2)
      }
    }
  }
}

func TestSet_Add(t *testing.T) {
  t.Run("Add unique numbers", testSet_AddFunc(New[int](), []int{1, 2, 3}, New[int](1, 2, 3)))
  t.Run("Add not unique numbers", testSet_AddFunc(New[int](3), []int{1, 2, 3, 2, 1}, New[int](1, 2, 3)))
}

func testSet_AddFunc(set *Set[int], items []int, exp *Set[int]) func(*testing.T) {
  return func(t *testing.T) {
    set.Add(items...)
    if !set.Equal(exp) {
      t.Errorf("Expected set to be equal: %v. Received: %v.", exp, set)
    }
  }
}

func TestSet_Size(t *testing.T) {
  set := New[int]()
  set.Add(1, 2, 3)

  t.Run("Empty set", testSet_SizeFunc(New[int](), 0))
  t.Run("Set with predefined items", testSet_SizeFunc(New[int](1, 2, 3, 4, 5), 5))
  t.Run("Set with added items", testSet_SizeFunc(set, 3))
}

func testSet_SizeFunc(set *Set[int], exp int) func(*testing.T) {
  return func(t *testing.T) {
    res := set.Size()
    if res != exp {
      t.Errorf("Expected set of size %v. Received size: %v.", exp, res)
    }
  }
}

func TestSet_Delete(t *testing.T) {
  t.Run("Normal removal 1", testSet_DeleteFunc(New[int](1, 2, 3), 2))
  t.Run("Normal removal 2", testSet_DeleteFunc(New[int](1, 2, 3), 8))
  t.Run("Empty set", testSet_DeleteFunc(New[int](), 5))
}

func testSet_DeleteFunc(set *Set[int], item int) func(*testing.T) {
  return func(t *testing.T) {
    set.Delete(item)
    if set.Has(item) {
      t.Errorf("Expected %v not to be in set after removal.", item)
    }
  }
}

func TestSet_Clear(t *testing.T) {
  t.Run("Normal set", testSet_ClearFunc(New[int](1, 2, 3)))
  t.Run("Empty set", testSet_ClearFunc(New[int]()))
}

func testSet_ClearFunc(set *Set[int]) func(*testing.T) {
  return func(t *testing.T) {
    set.Clear()
    size := set.Size()
    if size != 0 {
      t.Errorf("Expected empty set. Received set of size: %v.", size)
    }
  }
}

func TestSet_Copy(t *testing.T) {
  set := New[int](1, 2, 3)

  res := set.Copy()

  t.Run("Normal copy", testSet_CopyFunc(res, New[int](1, 2, 3)))
  set.Delete(2)
  t.Run("Deleting from original set after copy", testSet_CopyFunc(res, New[int](1, 2, 3)))
  res.Delete(3)
  t.Run("Deleting from copied set after copy", testSet_CopyFunc(set, New[int](1, 3)))
}

func testSet_CopyFunc(res *Set[int], exp *Set[int]) func(*testing.T) {
  return func(t *testing.T) {
    if !res.Equal(exp) {
      t.Errorf("Expected copied set to be equal: %v. Received: %v.", exp, res)
    }
  }
}

func TestSet_Union(t *testing.T) {
  t.Run("Normal sets", testSet_UnionFunc(New[int](1, 2, 3), New[int](2, 3, 4), New[int](1, 2, 3, 4)))
  t.Run("One empty set", testSet_UnionFunc(New[int](1, 2, 3), New[int](), New[int](1, 2, 3)))
}

func testSet_UnionFunc(set1 *Set[int], set2 *Set[int], exp *Set[int]) func(*testing.T) {
  return func(t *testing.T) {
    res := set1.Union(set2)
    if !res.Equal(exp) {
      t.Errorf("Expected union of %v and %v to be equal: %v. Received: %v.", set1, set2, exp, res)
    }
  }
}

func TestSet_Intersection(t *testing.T) {
  t.Run("Normal sets", testSet_IntersectionFunc(New[int](1, 2, 3), New[int](2, 3, 4), New[int](2, 3)))
  t.Run("One empty set", testSet_IntersectionFunc(New[int](1, 2, 3), New[int](), New[int]()))
}

func testSet_IntersectionFunc(set1 *Set[int], set2 *Set[int], exp *Set[int]) func(*testing.T) {
  return func(t *testing.T) {
    res := set1.Intersection(set2)
    if !res.Equal(exp) {
      t.Errorf("Expected intersection of %v and %v to be equal: %v. Received: %v.", set1, set2, exp, res)
    }
  }
}

func TestSet_Difference(t *testing.T) {
  t.Run("Normal sets", testSet_DifferenceFunc(New[int](1, 2, 3), New[int](2, 3, 4), New[int](1)))
  t.Run("Right empty set", testSet_DifferenceFunc(New[int](1, 2, 3), New[int](), New[int](1, 2, 3)))
  t.Run("Left empty set", testSet_DifferenceFunc(New[int](), New[int](1, 2, 3), New[int]()))
}

func testSet_DifferenceFunc(set1 *Set[int], set2 *Set[int], exp *Set[int]) func(*testing.T) {
  return func(t *testing.T) {
    res := set1.Difference(set2)
    if !res.Equal(exp) {
      t.Errorf("Expected difference between %v and %v to be equal: %v. Received: %v.", set1, set2, exp, res)
    }
  }
}
