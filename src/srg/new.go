package srg


func NewSrg(v, k, l, u int) *Srg {

	ss := make([][]bool, 1)
	ss[0] = []bool{false}

	srg := Srg{
		v:      v,
		k:      k,
		l:      l,
		u:      u,
		Matrix: make([]bool, v*(v+1)/2),
		filled: 0,
		toFill: v,
		SymSquare: ss,
	}

	return &srg
}

func NewFilledSrg(v, k, l, u int) *Srg {
	srg := NewSrg(v, k, l, u)

	row0 := make([]bool, v)
	for c := 1; c <= k; c++ {
		row0[c] = true
	}
	srg.Fill(row0)
	//fmt.Println(srg.PivotRow())
	
	row1 := make([]bool, v-1)
	for c := 1; c <= l; c++ {
		row1[c] = true
	}

	quotaForTrueUnderTrue := l
	quotaForTrueUnderFalse := k - l - 1 // 1 is from row0[1]
	for c := 1; c <= quotaForTrueUnderTrue; c++ {
		row1[c] = true
	}
	for c := k; c < k+quotaForTrueUnderFalse; c++ {
		row1[c] = true
	}
	srg.Fill(row1)
	//fmt.Println(srg.PivotRow())

	return srg
}