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

func NewQuestion(A *mat.Matrix, b, upperBound []int, ans *Answer) *Question {
	q:= newQuestion(A, b, upperBound)
	q.x = ans
	q.integrityCheck()
	return q
}


func NewQuestionWithDefaultAnswer(A *mat.Matrix, b, upperBound []int) *Question {
	q:= newQuestion(A, b, upperBound)
	q.x = NewAnswerDefault(A.C)
	q.integrityCheck()
	return q
}

func (q *Question) integrityCheck() {
	q.assert(q.x != nil, "answer is not initialized")
	q.assert(q.A.C == len(q.x.UnknownIndicies()), "answer and question out of sync")
}

func newQuestion(A *mat.Matrix, b, upperBound []int) *Question {
	return &Question{
		A : A,
		b : b,
		upperBound: upperBound,
		x: nil,
	}
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
	  
	s += fmt.Sprintf("      %v: upperBound\n", supperBound)
	s += fmt.Sprintf("Answer:\n")
	s += fmt.Sprintf("      %v\n" ,sx)
	s += fmt.Sprintf("      %v %3v\n" ,sloc, q.x.length)
	
	return s
}

func (q *Question) IsSolved() bool {
	if(!array.AllEqualTo(q.b, 0)) {return false}
	if(q.x.Unknown()) {return false}
	return true

}