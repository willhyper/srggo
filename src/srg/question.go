package srg


func (s *Srg) Question() ([]bool, []int, int){
	R := s.v - s.toFill //10 - 7 = 3
	C := s.toFill // 7
	
	A := make([]bool, R*C)
	i:=0
	// first row
	for c:=R;c<R+C;c++{
		A[i] = s.Matrix[c]
		i++
	}

	//rest rows
	for offset:=s.v - 1;i<R*C; offset+=offset-1{
		for c:=R;c<R+C;c++{
			A[i] = s.Matrix[c + offset]
			i++
		}
	}
	

	b := s.PivotQuotaLeft()

	pv := s.PivotRow()
	pvSum :=0
	for _, _pv := range pv{
		if _pv{
			pvSum+=1
		}
	}
	quota := s.k - pvSum

	return A, b, quota
}