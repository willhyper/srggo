package mat

import (
	"array"
)
func (m *Matrix) IndicesOfRow(row int) []int {
	ind := make([]int, m.C)
	offset := row * m.C
	for i := 0; i < m.C; i++ {
		ind[i] = offset + i
	}
	return ind
}

func (m *Matrix) IndicesOfCol(col int) []int {
	ind := make([]int, m.R)
	offset := 0
	for i := 0; i < m.R; i++ {
		ind[i] = offset + col
		offset += m.C
	}
	return ind
}

func (m *Matrix) EqualRows(row1, row2 int) bool {
	rind1 := m.IndicesOfRow(row1)
	rind2 := m.IndicesOfRow(row2)
	return array.Equal(m.arr, rind1, rind2)
}

func (m *Matrix) EqualCols(col1, col2 int) bool {
	cind1 := m.IndicesOfCol(col1)
	cind2 := m.IndicesOfCol(col2)
	return array.Equal(m.arr, cind1, cind2)
}

func (m *Matrix) Row(r int) []bool{
	return m.arr[r*m.C : r*m.C + m.C]
}

func (m *Matrix) Col(c int) []bool{
	cind := m.IndicesOfCol(c)
	col := make([]bool, m.R)
	for i, v := range cind{
		col[i] = m.arr[v]
	}
	return col
}

func (m *Matrix) IndexOfUniqueCols() []int {
	//assume Columns are sorted
	j:=0
	unique := append([]int(nil),j) // first col is always unique
	
	for i:=0 ; i<m.C ; i=j {
		for j = i+1 ; j<m.C ; j++ {
			if !m.EqualCols(i, j) {
				unique = append(unique, j)
				break
			}
		}
	}
	return unique
}

func (m *Matrix) SingleOutRow(r int) ([]bool, *Matrix) {
	a := m.arr
	R := m.R - 1
	C := m.C
	n := make([]bool, R*C)
	v := make([]bool, C)
	startExclude := r*C
	resume := startExclude + C
	for i:=0; i< startExclude;i++ {
		n[i] = a[i]
	}
	for i:=startExclude; i<resume ;i++ {
		v[i-startExclude] = a[i]
	}
	for i:=resume;i<len(a);i++ {
		n[i-C] = a[i]
	}
	return v, NewMatrix(n, R)
}

func (m *Matrix) SingleOutCol(c int) ([]bool, *Matrix) {
	a := m.arr
	R := m.R
	C := m.C // old matrix dimension
	n := make([]bool, R*C-R) //new matrix length
	v := make([]bool, R)

	for r:=0; r<R; r++{
		start := r*C
		exclude := start + c
		end := start + C

		for i:=start;i<exclude;i++{
			n[i-r]=a[i]
		}
		v[r]=a[exclude]
		for i:=exclude+1;i<end;i++{
			n[i-r-1] = a[i]
		}
	}
	
	return v, NewMatrix(n, R)
}
