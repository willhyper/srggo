package mat

import (
	"array"
	"fmt"
)

type Matrix struct {
	arr  []bool
	R, C int
}

func (m *Matrix) String() string {
	s := ""
	for r := 0; r < m.R; r++ {
		Mr := array.ToInt(m.Row(r))
		s += fmt.Sprintf("%v\n", Mr)
	}
	return s
}

func (m *Matrix) IsEmpty() bool {
	return len(m.arr) == 0
}