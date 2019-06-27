package srg

import "array"

// diagonal element is excluded
func (s *Srg) PivotRowIndices() []int {
	rowlen := s.v - s.toFill

	indices := make([]int, rowlen)
	indices[0] = rowlen
	for i := 1; i < rowlen; i++ {
		indices[i] = indices[i-1] + s.v - i
	}
	return indices
}

func (s *Srg) PivotRow() *[]bool {
	indices := s.PivotRowIndices()
	return array.GetValues(s.Matrix, indices)
}

func (s *Srg) PivotProduct() []int {
	pivotRowIndices := s.PivotRowIndices() //[4,13,21,28]
	N := len(pivotRowIndices)              // 4

	innerProd := make([]int, N)
	prodCount := 0

	indicesForProd := make([]int, N)
	for i := 0; i < N; i++ {
		indicesForProd[i] = i
	}
	innerProd[prodCount] = array.DotSrc(s.Matrix, pivotRowIndices, indicesForProd)
	prodCount++ // 1

	offset := s.v - 1 // 10-1=9
	for ; prodCount < N; prodCount++ {
		numOfOffset := N - prodCount // 4 - 1 = 3
		for i := N - 1; i >= N-numOfOffset; i-- {
			indicesForProd[i] += offset
		}
		for i := N - numOfOffset - 1; i >= 0; i-- {
			indicesForProd[i] += 1
		}
		innerProd[prodCount] = array.DotSrc(s.Matrix, pivotRowIndices, indicesForProd)
		offset--
	}

	return innerProd
}
