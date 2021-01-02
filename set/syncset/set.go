package syncset

import "sync"

type Set struct {
	m      *sync.Map
	length int
}

// NewSet
func NewSet(items ...interface{}) *Set {
	var s = &Set{m: new(sync.Map)}
	for i := 0; i < len(items); i++ {
		s.Add(items[i])
	}
	return s
}

// Add
func (s *Set) Add(k interface{}) bool {
	if !s.Contain(k) {
		return false
	}
	s.m.Store(k, struct{}{})
	s.length++
	return true
}

// Del
func (s *Set) Del(k interface{}) bool {
	if !s.Contain(k) {
		return false
	}
	s.m.Delete(k)
	s.length--
	return true
}

// Contain
func (s *Set) Contain(k interface{}) bool {
	var _, exist = s.m.Load(k)
	return exist
}

// Len
func (s *Set) Len() int {
	return s.length
}

// Equal 相等
func (s *Set) Equal(a *Set) bool {
	if s.Len() != a.Len() {
		return false
	}

	var flag bool
	a.m.Range(func(key, value interface{}) bool {
		if !s.Contain(key) {
			flag = false
			return false
		}
		flag = true
		return true
	})

	return flag
}

// IsSubset
func (s *Set) IsSubset(a *Set) bool {
	if a.Len() > s.Len() {
		return false
	}

	var flag bool
	a.Range(func(key, _ interface{}) bool {
		if !s.Contain(key) {
			flag = false
			return false
		}
		flag = true
		return true
	})

	return flag
}

// Range
func (s *Set) Range(f func(interface{}, interface{}) bool) {
	s.m.Range(f)
}

// Clear
func (s *Set) Clear() {
	s.m = new(sync.Map)
	s.length = 0
}

// 交集、并集、差集
// Intersection
func (s *Set) Intersection(a *Set) *Set {
	var x = NewSet()
	a.Range(func(k interface{}, _ interface{}) bool {
		if s.Contain(k) {
			x.Add(k)
		}
		return true
	})
	return x
}

// Union
func (s *Set) Union(a *Set) *Set {
	var x = NewSet()
	s.Range(func(k interface{}, _ interface{}) bool {
		x.Add(k)
		return true
	})
	a.Range(func(k interface{}, _ interface{}) bool {
		x.Add(k)
		return true
	})
	return x
}

// SymmetricDifference
func (s *Set) SymmetricDifference(a *Set) *Set {
	var x1 = s.Union(a)
	var x2 = s.Intersection(a)
	var ans = NewSet()
	x1.Range(func(k interface{}, _ interface{}) bool {
		if x2.Contain(k) {
			return true
		}
		ans.Add(k)
		return true
	})
	return ans
}
