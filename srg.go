package main

// 10,6,3,4
// [[0, 1, 1, 1, 1, 1, 1, 0, 0, 0],
//  [1, 0, 1, 1, 1, 0, 0, 1, 1, 0],
//  [1, 1, 0, 0, 0, 1, 1, 1, 1, 0],
//  [1, 1, 0, 0, 1, 1, 0, 1, 0, 1],
//  [1, 1, 0, 1, 0, 0, 1, 0, 1, 1],
//  [1, 0, 1, 1, 0, 0, 1, 1, 0, 1],
//  [1, 0, 1, 0, 1, 1, 0, 0, 1, 1],
//  [0, 1, 1, 1, 0, 1, 0, 0, 1, 1],
//  [0, 1, 1, 0, 1, 0, 1, 1, 0, 1],
//  [0, 0, 0, 1, 1, 1, 1, 1, 1, 0]]

import "fmt"
import "strconv"

type Srg struct {
	v, k, l, u int
	Matrix     []bool
	filled     int
	toFill     int
	SymSquare  [][]bool
}

var _map = map[bool]int{false: 0, true: 1}

func NewSrg(v, k, l, u int) *Srg {

	ss := make([][]bool, 1)
	ss[0] = []bool{false}

	srg := Srg{
		v:      v,
		k:      k,
		l:      l,
		u:      u,
		Matrix: make([]bool, v*(v+1)/2),
		filled: 0,
		toFill: v,
		SymSquare: ss,
	}

	return &srg
}

func NewFilledSrg(v, k, l, u int) *Srg {
	srg := NewSrg(v, k, l, u)

	row0 := make([]bool, v)
	for c := 1; c <= k; c++ {
		row0[c] = true
	}
	srg.Fill(row0)
	//fmt.Println(srg.PivotRow())
	
	row1 := make([]bool, v-1)
	for c := 1; c <= l; c++ {
		row1[c] = true
	}

	quotaForTrueUnderTrue := l
	quotaForTrueUnderFalse := k - l - 1 // 1 is from row0[1]
	for c := 1; c <= quotaForTrueUnderTrue; c++ {
		row1[c] = true
	}
	for c := k; c < k+quotaForTrueUnderFalse; c++ {
		row1[c] = true
	}
	srg.Fill(row1)
	//fmt.Println(srg.PivotRow())

	return srg
}

// diagonal element is excluded
func (srg *Srg) PivotRow() ([]bool, []int){
	rowlen := srg.v - srg.toFill

	indices := make([]int, rowlen)
	pivotRow := make([]bool, rowlen)
	
	indices[0] = rowlen
	for i:=1;i<rowlen;i++{
		indices[i]=indices[i-1] + srg.v - i
	}

	for i, v := range indices{
		pivotRow[i]= srg.Matrix[v]
	}
	return pivotRow, indices

}


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

func (srg *Srg) Println() {
	mi := make([]int, len(srg.Matrix))
	for i, v := range srg.Matrix {
		mi[i] = _map[v]
	}

	paddingFormat := "%" + strconv.Itoa(srg.v*2+2) + "v\n"
	used := 0
	for c := srg.v; c > 0; c-- {
		str := fmt.Sprint(mi[used : used+c])
		fmt.Printf(paddingFormat, str)
		used += c
	}

}
func main() {
	srg := NewFilledSrg(10, 6, 3, 4)
	srg.Println()

}
