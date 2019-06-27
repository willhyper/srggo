package srg

import "array"

// diagonal element is excluded
func (s *Srg) PivotRowIndices() []int{
	rowlen := s.v - s.toFill

	indices := make([]int, rowlen)
	indices[0] = rowlen
	for i:=1;i<rowlen;i++{
		indices[i]=indices[i-1] + s.v - i
	}
	return indices
}

func (s *Srg) PivotRow() *[]bool{
	indices := s.PivotRowIndices()
	return array.GetValues(s.Matrix, indices)
}