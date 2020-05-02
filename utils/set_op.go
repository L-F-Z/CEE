package utils

import (
	"math"
)

// SetUnique delete duplicate elements in a slice
func SetUnique(a []int64) []int64 {
	m := make(map[int64]bool)
	c := make([]int64, 0)
	for _, elem := range a {
		if _, exists := m[elem]; !exists {
			m[elem] = true
			c = append(c, elem)
		}
	}
	return c
}

// SetUnion returns a union of a and b
func SetUnion(a, b []int64) []int64 {
	m := make(map[int64]bool)
	c := make([]int64, 0)
	for _, elem := range a {
		m[elem] = true
	}
	for _, elem := range b {
		m[elem] = true
	}
	for key := range m {
		c = append(c, key)
	}
	return c
}

// SetIntersect returns a intersect of a and b
func SetIntersect(a, b []int64) []int64 {
	m := make(map[int64]bool)
	c := make([]int64, 0)
	for _, elem := range a {
		m[elem] = true
	}
	for _, elem := range b {
		if _, exists := m[elem]; exists {
			c = append(c, elem)
		}
	}
	return c
}

// SetMinus returns a minus of a and b
func SetMinus(a, b []int64) []int64 {
	m := make(map[int64]bool)
	c := make([]int64, 0)
	for _, elem := range a {
		m[elem] = true
	}
	for _, elem := range b {
		m[elem] = false
	}
	for key, elem := range m {
		if elem {
			c = append(c, key)
		}
	}
	return c
}

// PowerSet return all subsets of a
func PowerSet(a []int64) [][]int64 {
	size := int64(math.Pow(2, float64(len(a))))
	c := make([][]int64, 0, size)
	var i int64
	for i < size {
		var subSet []int64
		for j, elem := range a {
			if i&(1<<uint(j)) > 0 {
				subSet = append(subSet, elem)
			}
		}
		c = append(c, subSet)
		i++
	}
	return c
}

// SameSet judges if two sets are the same
func SameSet(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[int64]bool)
	for _, val := range a {
		m[val] = true
	}
	for _, val := range b {
		if _, exists := m[val]; !exists {
			return false
		}
	}
	return true
}
