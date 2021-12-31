// Package set provides universal set implementation using go1.18beta1 generics.
package set

import "fmt"

// Set is implemented using Go's map.
// It maps comparable type to empty structs which use 0 memory compared to other types.
type Set[K comparable] map[K]struct{}

// New creates and returns a pointer to a new set.
// It can be called with elems, which are elements to be added.
func New[K comparable](elems ...K) *Set[K] {
	set := make(Set[K])
	for _, e := range elems {
		set[e] = struct{}{}
	}
	return &set
}

// ToSlice returns slice converted from set.
func (set *Set[K]) ToSlice() []K {
	res := make([]K, 0, set.Size())
	for e := range *set {
		res = append(res, e)
	}
	return res
}

// String returns string representation of set. Like:
// set[1 2 3] represents set with elements: 1, 2 and 3
// set[] represents empty set
func (set *Set[K]) String() string {
	str := "set"
	slice := set.ToSlice()
	str += fmt.Sprint(slice)
	return str
}

// Has reports whether e is in the set.
func (set *Set[K]) Has(e K) bool {
	_, ok := (*set)[e]
	return ok
}

// Add adds elems to the set.
func (set *Set[K]) Add(elems ...K) {
	for _, e := range elems {
		(*set)[e] = struct{}{}
	}
}

// Size returns actual number of items in the set.
func (set *Set[K]) Size() int {
	return len(*set)
}

// Delete removes item e from the set.
// It does not inform whether item e was in set. It should be checked by a programmer calling Has method first.
func (set *Set[K]) Delete(e K) {
	delete(*set, e)
}

// Clear clears whole set.
func (set *Set[K]) Clear() {
	// Deleting from map in a loop is optimized by compiler since go1.11.
	for e := range *set {
		delete(*set, e)
	}
}

// Copy returns copy of the set.
func (set *Set[K]) Copy() *Set[K] {
	res := New[K]()
	for e := range *set {
		res.Add(e)
	}
	return res
}

// Equal checks equality of set1 and set2 in terms of equality of elements into them.
func (set1 *Set[K]) Equal(set2 *Set[K]) bool {
	if set1.Size() != set2.Size() {
		return false
	}

	for e := range *set1 {
		if !set2.Has(e) {
			return false
		}
	}

	for e := range *set2 {
		if !set1.Has(e) {
			return false
		}
	}
	return true
}

// Union returns a new set which is an union of set1 and set2.
func (set1 *Set[K]) Union(set2 *Set[K]) *Set[K] {
	res := New[K]()
	for e := range *set1 {
		res.Add(e)
	}

	for e := range *set2 {
		res.Add(e)
	}
	return res
}

// Intersection returns a new set which is an intersection of set1 and set2.
func (set1 *Set[K]) Intersection(set2 *Set[K]) *Set[K] {
	res := New[K]()
	if set1.Size() > set2.Size() {
		set1, set2 = set2, set1 // Swap to iterate over shorter set
	}

	for e := range *set1 {
		if set2.Has(e) {
			res.Add(e)
		}
	}
	return res
}

// Difference returns a new set which is an difference between set1 and set2.
func (set1 *Set[K]) Difference(set2 *Set[K]) *Set[K] {
	res := New[K]()
	for e := range *set1 {
		if !set2.Has(e) {
			res.Add(e)
		}
	}
	return res
}
