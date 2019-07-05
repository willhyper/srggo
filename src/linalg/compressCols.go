package linalg

import (
	"array"
)

func (q *Question) CompressCols() *Question {
	m := q.A

	cols := m.IndexOfUniqueCols()

	mReduced := m.Cols(cols)

	upperBound := array.Diff(cols)
	last := m.C - array.SumInt(upperBound)
	upperBound = append(upperBound, last)

	return NewQuestion(mReduced, q.b, upperBound)
}