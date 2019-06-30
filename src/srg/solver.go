package srg

import(
	"linalg"
)

func Solve(q *Srg) (<-chan []bool){
	answerCh := make(chan []bool)

	if q.toFill == 0{
		panic("question is solved")
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
			for x := range linalg.Solver(A, b, quota){
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