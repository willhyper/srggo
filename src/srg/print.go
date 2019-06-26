package srg

import "fmt"
import "strconv"


var _map = map[bool]int{false: 0, true: 1}

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
