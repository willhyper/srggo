package srg

import (
	"array"
	"mat"
	"linalg"
)

func (s *Srg) ToFill() int {
	return s.toFill
}
func (s *Srg) Question() (*linalg.Question) {
	R := s.v - s.toFill //10 - 7 = 3
	C := s.toFill - 1 // 7 - 1 = 6 
	//-1 because first element in answer row is always false. 
	// no bother formulate question for answer that is known

	A := make([]bool, R*C)
	i:=0
	// first row
	for c:=R+1;c<=R+C;c++{
		A[i] = s.Matrix[c]
		i++
	}
	//rest rows
	cumoffset := 0
	for offset:=s.v - 1;i<R*C; offset--{
		cumoffset += offset
		for c:=R+1;c<=R+C;c++{
			A[i] = s.Matrix[c + cumoffset]
			i++
		}
	}
	b := s.PivotQuotaLeft()
	quota := s.k - array.SumBool(s.PivotRow())

	// k condition 
	ones := make([]bool, C)
	upperBound := make([]int, C)
	for c:=0 ; c<C ; c++{
		ones[c] = true
		upperBound[c] = 1
	}
	Ak := append(A, ones...)
	bk := append(b, quota)	
	Matk := mat.NewMatrix(Ak, len(bk))

	return linalg.NewQuestionWithDefaultAnswer(Matk, bk, upperBound)
}