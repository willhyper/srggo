package linalg

import (
	"array"
)

func Fork(q *Question, qchan chan *Question) {

	q.assert(q.x.Unknown(), "question is solved. no fork required")

	ci := array.FirstMinIndex(q.upperBound)
	col, Anew := q.A.SingleOutCol(ci)
	upperBoundNew, upperBoundAtci := array.SingleOutElement(q.upperBound, ci)
	for ans := 0; ans <= upperBoundAtci; ans++ {
		anscol := array.Multiply(col, ans)
		bnew := array.SubstractVector(q.b, anscol)

		if array.AllNonnegative(bnew){
			ansnew := q.x.Copy()
			ansnew.FillUnknown(ans, ci)
	
			qnew := NewQuestion(Anew, bnew, upperBoundNew, ansnew)
			qchan <- qnew
		}
	}
}
