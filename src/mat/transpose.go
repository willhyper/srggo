package mat

func (m *Matrix) Transpose() *Matrix {

	C := m.C
	R := m.R

	mtarr := make([]bool, R*C)

	for i := 0; i < R*C; i++ {
		p := (i%R)*C + i/R
		mtarr[i] = m.arr[p]
	}
	return NewMatrix(mtarr, C)
}
