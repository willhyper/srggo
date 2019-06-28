package array

var _map = map[bool]int{false: 0, true: 1}


func ToInt(a []bool) *[]int{
	mi := make([]int, len(a))
	for i, v := range a {
		mi[i] = _map[v]
	}
	return &mi
}