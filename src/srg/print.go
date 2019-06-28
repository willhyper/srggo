package srg


import(
	"fmt"
	"strconv"
	"array"
)

func (srg *Srg) Println() {
	paddingFormat := "%" + strconv.Itoa(srg.v*2+2) + "v\n"
	used := 0
	for c := srg.v; c > 0; c-- {
		str := array.Sprint(srg.Matrix[used : used+c])
		fmt.Printf(paddingFormat, str)
		used += c
	}
}
