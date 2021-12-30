package ds

import "fmt"

type Set[K comparable] map[K]struct{} // empty struct use 0 bytes of memory

func NewSet[K comparable](elems ...K) *Set[K] {
	set := make(Set[K])
	for _, e := range elems {
		set[e] = struct{}{}
	}
	return &set
}

func (set *Set[K]) ToSlice() []K {
	res := make([]K, 0, set.Size())
	for e := range *set {
		res = append(res, e)
	}
	return res
}

func (set *Set[K]) String() string {
	str := "set"
	slice := set.ToSlice()
	str += fmt.Sprint(slice)
	return str
}

func (set *Set[K]) Has(k K) bool {
	_, ok := (*set)[k]
	return ok
}

func (set *Set[K]) Add(elems ...K) {
	for _, e := range elems {
		(*set)[e] = struct{}{}
	}
}

func (set *Set[K]) Size() int {
	return len(*set)
}

func (set *Set[K]) Delete(k K) {
	delete(*set, k)
}

func (set *Set[K]) Clear() {
	for e := range *set { // optimized by compiler since go1.11
		delete(*set, e)
	}
}

func (set *Set[K]) Copy() *Set[K] {
	res := NewSet[K]()
	for e := range *set {
		res.Add(e)
	}
	return res
}

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

func (set1 *Set[K]) Union(set2 *Set[K]) *Set[K] {
	res := NewSet[K]()
	for e := range *set1 {
		res.Add(e)
	}

	for e := range *set2 {
		res.Add(e)
	}
	return res
}

func (set1 *Set[K]) Intersection(set2 *Set[K]) *Set[K] {
	res := NewSet[K]()
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

func (set1 *Set[K]) Difference(set2 *Set[K]) *Set[K] {
	res := NewSet[K]()
	for e := range *set1 {
		if !set2.Has(e) {
			res.Add(e)
		}
	}
	return res
}
