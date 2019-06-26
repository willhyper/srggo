package srg

import "fmt"

func (srg *Srg) Fill(row []bool) {
	if lenrow := len(row); lenrow != srg.toFill {
		panic(fmt.Sprintf("Fill row of wrong length. Expecting len %v but got %v\n", srg.toFill, lenrow))
	}
	if row[0] == true {
		panic(fmt.Sprintf("1st element of the row shall always be false"))
	}

	cc := 0
	for c := srg.filled; c < srg.filled+srg.toFill; c++ {
		srg.Matrix[c] = row[cc]
		cc++
	}
	srg.filled += srg.toFill
	srg.toFill -= 1

	// dim := len(srg.SymSquare)
	// update SymSquare content from PivotRow, which diagonal element is excluded.
	// size from dim x dim to (dim+1) x (dim+1)
	pv, _ := srg.PivotRow()
	for ci, cv := range srg.SymSquare{
		srg.SymSquare[ci] = append(cv, pv[ci])
	}
	pv = append(pv, false) // diagnonal element is always false
	srg.SymSquare = append(srg.SymSquare, pv)
}