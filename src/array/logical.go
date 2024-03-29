package array

func Equal(src []bool, indices1, indices2[]int) bool{
	N:=len(indices1)
	if N!=len(indices2){
		panic("len(indices1)!=len(indices2)")
	}
	for i:=0;i<N;i++{
		if src[indices1[i]] != src[indices2[i]]{
			return false
		} 
	}
	return true
}

// find the nth occurrence index of val in array arr
// if not found, return -1
func FindIntIndex(arr []int, val, nth int) int{
	occurence := 0
	for i, v := range arr {
		if v==val {
			if occurence == nth {
				return i
			} else {
				occurence++
			}
		}
	}
	return -1
}

// find all occurrence indices of val in array arr
func FindInt(arr []int, val int) []int {
	ind := make([]int, len(arr))
	occurence := 0
	for i, v := range arr {
		if v==val {
			ind[occurence] = i
			occurence++
		}
	}
	return ind[:occurence]
}

// return the index of first min value in array a
func FirstMinIndex(a []int) int {
	if len(a) == 0{panic("empty array")}
	min, imin := a[0], 0
	for i:=1;i<len(a);i++ {
		if a[i] < min {
			min, imin = a[i], i
		}
	}
	return imin
}

func Min(a ...int) int {
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func AllEqualTo(a []int, val int) bool {
	for _, v := range a {
		if v != val {
			return false
		}
	}
	return true
}

func AllNonnegative(a []int) bool {
	for _, v := range a {
		if v<0 {
			return false
		}
	}
	return true
}

