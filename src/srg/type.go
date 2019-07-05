package srg

import(
	"fmt"
	"strconv"
	"array"
)

type Srg struct {
	v, k, l, u int
	Matrix     []bool
	filled     int
	toFill     int
}


func (srg *Srg) String() string {
	s := ""
	paddingFormat := "%" + strconv.Itoa(srg.v*2+2) + "v\n"
	used := 0
	for c := srg.v; c > srg.toFill; c-- {
		row := array.ToInt(srg.Matrix[used : used+c])
		rowstr := fmt.Sprintf("%v",row)
		s += fmt.Sprintf(paddingFormat, rowstr)
		used += c
	}
	return s
}
