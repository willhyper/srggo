package linalg

import (
	"mat"
	"fmt"
	"array"
)

type Question struct{
	A *mat.Matrix
	b []int
	condition []int
	upperBound []int
}

func NewQuestion(A *mat.Matrix, b, condition, upperBound []int) *Question {

	q := Question{
		A : A,
		b : b,
		condition : condition,
		upperBound: upperBound,
	}

	// check
	for e:=0; e<len(condition); e++ {
		q.assert(condition[e] >= upperBound[e], "condition < upperBound")
		q.assert(upperBound[e] >=0, "upperBound < 0")
	}
	for e:=0; e<len(b) ; e++{
		q.assert(b[e] >= 0, "b < 0")
	}

	return &q
}
func (q *Question) AnswerBool(ansInt []int) []bool {
	
	//check
	for e:=0; e< len(q.upperBound); e++{
		q.assert(ansInt[e]<= q.upperBound[e], fmt.Sprintf("upperBound < ansInt %v", ansInt))
	}

	l := array.SumInt(q.condition) // [2,2,2,1] => 7
	ansBool := make([]bool, l)

	p:=0
	for i, v := range q.condition {
		numOfTrue := ansInt[i]
		for j:=0; j<numOfTrue;j++{
			ansBool[p + j] = true		
		}
		p += v
	}
	if(p != l) {panic("design error")}

	return ansBool
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
		Ar := *array.ToInt(q.A.Row(r))
		s += fmt.Sprintf("%3v | %v\n", q.b[r], Ar)
	}
	s += fmt.Sprintf("      %v : condition\n", q.condition)
	s += fmt.Sprintf("      %v : upperBound\n", q.upperBound)
	return s
}