package mat

func (m *Matrix) Cols(cind []int) *Matrix {
	R := m.R
	C := m.C
	CNew := len(cind)

	arr := make([]bool, R*CNew)

	i:=0
	for r:=0; r<R; r++{
		for _, c:= range(cind){
			arr[i] = m.arr[r*C + c]
			i++
		}
	}
	return NewMatrix(arr, R)
}

func (m *Matrix) Rows(rind []int) *Matrix {
	//R := m.R
	C := m.C
	RNew := len(rind)

	arr := make([]bool, RNew*C)

	i:=0
	for _, r:= range rind {
		start := r*C
		end := start + C
		for j:=start; j<end;j++{
			arr[i] = m.arr[j]
			i++
		}
	}
	return NewMatrix(arr, RNew)
}