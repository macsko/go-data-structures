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
	set1 := New[int](1, 2, 3)
	set2 := New[int](3, 4)

	exp := []int{1, 2, 3}
	res := set1.ToSlice()
	if !slicesContentEqual(exp, res) {
		t.Errorf("Expected converted slice to be equal: %v. Received: %v.", exp, res)
	}

	exp = []int{3, 4}
	res = set2.ToSlice()
	if !slicesContentEqual(exp, res) {
		t.Errorf("Expected converted slice to be equal: %v. Received: %v.", exp, res)
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
	set := New[int](1, 2, 3)

	res := set.Has(2)
	if res == false {
		t.Error("Expected true for set[1 2 3]. Received:", res)
	}

	res = set.Has(9)
	if res == true {
		t.Error("Expected false for set[1 2 3]. Received:", res)
	}
}

func TestSet_Equal(t *testing.T) {
	set1 := New[int](1, 2, 3)
	set2 := New[int](1, 2, 3)
	if !set1.Equal(set2) {
		t.Error("Expected set1 to be equal to set2.")
	}
	set3 := New[int](1, 2)
	if set1.Equal(set3) {
		t.Error("Expected set1 not to be equal to set3.")
	}
}

func TestSet_Add(t *testing.T) {
	set := New[int]()
	set.Add(1, 2, 3, 2, 1)

	exp := New[int](1, 2, 3)
	if !set.Equal(exp) {
		t.Errorf("Expected set to be equal: %v. Received: %v.", exp, set)
	}
}

func TestSet_Size(t *testing.T) {
	set := New[int]()

	exp := 0
	res := set.Size()
	if res != exp {
		t.Errorf("Expected set of size %v. Received size: %v.", exp, res)
	}

	exp = 3
	set.Add(1, 2, 3)
	res = set.Size()
	if res != exp {
		t.Errorf("Expected set of size %v. Received size: %v.", exp, res)
	}
}

func TestSet_Delete(t *testing.T) {
	set := New[int](1, 2, 3)

	set.Delete(2)
	if set.Has(2) {
		t.Error("Expected 2 not to be in set.")
	}

	set.Delete(8)
	if set.Has(8) {
		t.Error("Expected 8 not to be in set.")
	}
}

func TestSet_Clear(t *testing.T) {
	set := New[int](1, 2, 3)

	set.Clear()
	size := set.Size()
	if size != 0 {
		t.Error("Expected empty set. Received set of size:", size)
	}
}

func TestSet_Copy(t *testing.T) {
	set := New[int](1, 2, 3)

	exp := New[int](1, 2, 3)
	res := set.Copy()
	if !res.Equal(exp) {
		t.Errorf("Expected copied set to be equal: %v. Received: %v.", exp, res)
	}

	set.Delete(2)
	if !res.Equal(exp) {
		t.Errorf("Expected copied set to be equal: %v. Received: %v.", exp, res)
	}
}

func TestSet_Union(t *testing.T) {
	set1 := New[int](1, 2, 3)
	set2 := New[int](2, 3, 4)
	emptySet := New[int]()

	exp := New[int](1, 2, 3, 4)
	res := set1.Union(set2)
	if !res.Equal(exp) {
		t.Errorf("Expected union to be equal: %v. Received: %v.", exp, res)
	}

	exp = set1
	res = set1.Union(emptySet)
	if !res.Equal(exp) {
		t.Errorf("Expected union to be equal: %v. Received: %v.", exp, res)
	}
}

func TestSet_Intersection(t *testing.T) {
	set1 := New[int](1, 2, 3)
	set2 := New[int](2, 3, 4)
	emptySet := New[int]()

	exp := New[int](2, 3)
	res := set1.Intersection(set2)
	if !res.Equal(exp) {
		t.Errorf("Expected intersection to be equal: %v. Received: %v.", exp, res)
	}

	exp = emptySet
	res = set1.Intersection(emptySet)
	if !res.Equal(exp) {
		t.Errorf("Expected intersection to be equal: %v. Received: %v.", exp, res)
	}
}

func TestSet_Difference(t *testing.T) {
	set1 := New[int](1, 2, 3)
	set2 := New[int](2, 3, 4)
	emptySet := New[int]()

	exp := New[int](1)
	res := set1.Difference(set2)
	if !res.Equal(exp) {
		t.Errorf("Expected difference to be equal: %v. Received: %v.", exp, res)
	}

	exp = set1
	res = set1.Difference(emptySet)
	if !res.Equal(exp) {
		t.Errorf("Expected difference to be equal: %v. Received: %v.", exp, res)
	}
}
