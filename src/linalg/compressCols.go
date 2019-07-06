package linalg

import (
	"array"
)

func (q *Question) CompressCols() {
	m := q.A

	cols := m.IndexOfUniqueCols() // [0,2,4,6]

	mReduced := m.Cols(cols)

	upperBound := array.Diff(cols) // [2,2,2]
	last := m.C - array.SumInt(upperBound) // 7-6 = 1
	upperBound = append(upperBound, last) // [2,2,2,1]

	arrNew := array.Ones(mReduced.C, UNKNOWN)
	ans := NewAnswer(arrNew, cols, q.x.length)

	q.A = mReduced
	q.upperBound = upperBound
	q.x = ans
	//q.b = q.b // no change
}