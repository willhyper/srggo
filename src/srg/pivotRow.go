package srg

// diagonal element is excluded
func (srg *Srg) PivotRow() ([]bool, []int){
	rowlen := srg.v - srg.toFill

	indices := make([]int, rowlen)
	pivotRow := make([]bool, rowlen)
	
	indices[0] = rowlen
	for i:=1;i<rowlen;i++{
		indices[i]=indices[i-1] + srg.v - i
	}

	for i, v := range indices{
		pivotRow[i]= srg.Matrix[v]
	}
	return pivotRow, indices

}
