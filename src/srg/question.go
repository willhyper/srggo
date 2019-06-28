package srg

import "array"

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

	quota := s.k - array.Sum(s.PivotRow())

	return A, b, quota
}