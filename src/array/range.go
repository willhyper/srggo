package array

import (
)
// return a range [0, N) without compliments
func CompRange(N int, compliments []int) []int {
	// assume compliments are sorted
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i
	}

	if len(compliments) == 0 {
		return a
	}

	for i := len(compliments) - 1; i >= 0; i-- {
		c := compliments[i]
		a = append(a[:c], a[c+1:]...)
	}
	return a
}

// return a[indices] as []int
func Select(a []int, indices []int) []int {
	b := make([]int, len(indices))
	for i, v := range indices {
		b[i] = a[v]
	}
	return b
}

func Reverse(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func Unique(a []int) []int {
    keys := make(map[int]bool)
    list := []int{} 
    for _, entry := range a {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}