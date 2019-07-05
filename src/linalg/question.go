package linalg

import (
	"mat"
	"fmt"
	"array"
)

type Question struct{
	A *mat.Matrix
	b []int
	upperBound []int
	x *Answer
}

func NewQuestion(A *mat.Matrix, b, upperBound []int) *Question {

	q := Question{
		A : A,
		b : b,
		upperBound: upperBound,
		x: NewAnswerDefault(len(upperBound)),
	}

	// check
	for e:=0; e<len(upperBound); e++ {
		q.assert(upperBound[e] >=0, "upperBound < 0")
	}
	for e:=0; e<len(b) ; e++{
		q.assert(b[e] >= 0, "b < 0")
	}

	return &q
}


func (q *Question) assert(condition bool, errorMessage string){
	if !condition {
		fmt.Println(q)
		panic(errorMessage)
	}
}

func (q *Question) String() string {
	s := ""
	for r:=0;r<q.A.R;r++{
		Ar := array.ToInt(q.A.Row(r))
		Arstr := array.ToString("%3v", Ar)
		s += fmt.Sprintf("%3v | %v\n", q.b[r], Arstr)
	}
	supperBound:= array.ToString("%3v", q.upperBound)
	sx         := array.ToString("%3v", q.x.arr)
	sloc       := array.ToString("%3v", q.x.location)
	  
	s += fmt.Sprintf("      %v    : upperBound\n", supperBound)
	s += fmt.Sprintf("      %v    : answer\n" ,sx)
	s += fmt.Sprintf("      %v %3v: location\n" ,sloc, q.x.length)
	
	return s
}