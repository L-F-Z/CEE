package utils

import ()

// SliceInSlice checks if b in a
func SliceInSlice(a [][]int64, b []int64) bool {
	m := make(map[int64]bool)
	for _, e := range b {
		m[e] = true
	}
	lenb := len(b)
	for _, s := range a {
		if len(s) != lenb {
			continue
		}
		same := true
		for _, e := range s {
			if _, exists := m[e]; !exists {
				same = false
			}
		}
		if same {
			return true
		}
	}
	return false
}

// SliceContainSlice function
func SliceContainSlice(a [][]int64, b []int64) []int64 {
	for _, s := range a {
		if len(b) >= len(s) {
			continue
		}
		m := make(map[int64]bool)
		for _, e := range s {
			m[e] = true
		}
		contain := true
		for _, e := range b {
			if _, exists := m[e]; !exists {
				contain = false
			}
		}
		if contain {
			return s
		}
	}
	return nil
}