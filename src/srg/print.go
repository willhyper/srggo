package srg


import(
	"fmt"
	"strconv"
	"array"
)

func (srg *Srg) Println() {
	paddingFormat := "%" + strconv.Itoa(srg.v*2+2) + "v\n"
	used := 0
	for c := srg.v; c > srg.toFill; c-- {
		str := array.Sprint(srg.Matrix[used : used+c])
		fmt.Printf(paddingFormat, str)
		used += c
	}
}


func (s *Srg) PrintQuestion(){
	A, b, quota := s.Question()
	R := s.v - s.toFill //10 - 7 = 3
	C := s.toFill - 1 // 7 - 1 = 6 

	for r:=0;r<R;r++{
		Ar := *array.ToInt(A[r*C: r*C+C])
		fmt.Printf("%3v | %v\n", b[r], Ar)
	}
	fmt.Println("quota", quota)	
}