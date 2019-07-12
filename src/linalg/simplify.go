package linalg

import (
	"array"
	//"fmt"
)

func Simplify(q *Question) {
	compressCols(q) // this enforces 1's moved ahead
	zeroUpperBound(q)
}

func zeroUpperBound(q *Question) {
	rzeros := array.FindInt(q.b, 0)
	if len(rzeros)==0 {return}

	// construct czeros
	var czeros []int
	for _, r := range rzeros {
		Ar := q.A.Row(r)
		for c, v := range Ar {
			if v {
				czeros = append(czeros, c)
			}
		}
	}// czeros is ready [1,2,5]
	if len(czeros)==0 {return}
	czeros = array.Unique(czeros)

	ckeeps :=array.CompRange(q.A.C, czeros) // 8, [1,2,5] => [0,3,4,6,7]
	rkeeps :=array.CompRange(q.A.R, rzeros)

	bkeeps :=array.Select(q.b, rkeeps)
	upperBoundKeep := array.Select(q.upperBound, ckeeps)
	Akeeps := q.A.Cols(ckeeps).Rows(rkeeps)
	
	// update q
	array.Reverse(czeros)
	for _, c := range czeros {
		q.x.FillUnknown(0, c) // Fill zero backwards so smaller indices are unchanged
	}
	q.A = Akeeps
	q.b = bkeeps
	q.upperBound = upperBoundKeep
}

func compressCols(q *Question) {
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