package mat

import "array"

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