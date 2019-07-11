package linalg

import (
	"array"
)

func Simplify(q *Question) {
	q.CompressCols()	
}

func (q *Question) CompressCols() {
	m := q.A
	if m.C < 2 {return} // no bother compressing

	colsInQuestionSpace := m.IndexOfUniqueCols() // [0,2,4,6]
	// new matrix
	mReduced := m.Cols(colsInQuestionSpace)

	// new upperBound
	numOfBound  := len(colsInQuestionSpace)
	upperBoundNew := make([]int, numOfBound)
	locs := append(colsInQuestionSpace, m.C) // [0,2,4,6,7] = append([0,2,4,6],7)
	for i:=0; i< numOfBound;i++ { // i = 0,1,2,3
		sumStart := locs[i]
		sumEnd := locs[i+1] // excluded
		sum := 0
		for j:=sumStart;j<sumEnd;j++{
			sum += q.upperBound[j]
		}
		upperBoundNew[i] = sum // [2,2,2,1]
	}


	// new answer
	unknownIndicies := q.x.UnknownIndicies()
	colsInAnswerSpace:=make([]int, len(colsInQuestionSpace)) // the indices to turn from unknown to known
	for i, v:= range colsInQuestionSpace {
		colsInAnswerSpace[i] = unknownIndicies[v]
	}
	colsToKeepInAnswerSpace := array.Merge(q.x.KnownIndicies(), colsInAnswerSpace)
	ans := q.x.CopySelect(colsToKeepInAnswerSpace)
	
	// update q in place
	q.A = mReduced
	q.upperBound = upperBoundNew
	q.x = ans
	//q.b = q.b // no change
}