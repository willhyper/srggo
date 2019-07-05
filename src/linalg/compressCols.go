package linalg

import (
	"array"
)

func (q *Question) CompressCols() *Question {
	m := q.A

	cols := m.IndexOfUniqueCols()

	mReduced := m.Cols(cols)

	condition := array.Diff(cols)
	last := m.C - array.SumInt(condition)
	condition = append(condition, last)

	return NewQuestion(mReduced, q.b, condition, condition)
}