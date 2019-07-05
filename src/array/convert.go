package array

import (
	"fmt"
)
var _map = map[bool]int{false: 0, true: 1}


func ToInt(a []bool) []int{
	mi := make([]int, len(a))
	for i, v := range a {
		mi[i] = _map[v]
	}
	return mi
}

func ToString(format string, a []int) string {
	s:=""
	for i:=0;i<len(a);i++ {
		s += fmt.Sprintf(format, a[i])
	}
	return "[" + s + "]"
}
