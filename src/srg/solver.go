package srg

import(
	"linalg"
	"mat"
)

func Solve(q *Srg) (<-chan []bool){
	answerCh := make(chan []bool)

	if q.toFill == 0{
		panic("question is solved. Should not invoke")
	}else if q.toFill == 1{
		go func(){
			defer close(answerCh)
			answer := []bool{false}
			answerCh <- answer
		}()
	}else{
		A, b, quota := q.Question()
		go func(){
			defer close(answerCh)

			C := q.toFill - 1
			ones := make([]bool, C)
			for c:=0 ; c<C ; c++{
				ones[c] = true
			}
			Ak := append(A, ones...)
			bk := append(b, quota)	
			Matk := mat.NewMatrix(Ak, len(bk))

			for x := range linalg.Solver(Matk, bk){
				answer := append([]bool{false}, x...)
				answerCh <- answer
			}
		}()
	}

	return answerCh
}

func (s *Srg) IsSolved() bool{
	return s.toFill == 0
}