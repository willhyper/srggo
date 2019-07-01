package srg

import(
	"linalg"
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
		A, b, _ := q.Question()
		go func(){
			defer close(answerCh)

			for x := range linalg.Solver(A, b){
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