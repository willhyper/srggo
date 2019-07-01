package srg

import(
	"linalg"
)

func Solve(s *Srg) (<-chan []bool){
	answerCh := make(chan []bool)

	if s.toFill == 0{
		panic("question is solved. Should not invoke")
	}else if s.toFill == 1{
		go func(){
			defer close(answerCh)
			answer := []bool{false}
			answerCh <- answer
		}()
	}else{
		q := s.Question()
		go func(){
			defer close(answerCh)

			for xInt := range linalg.Solver(q){
				xBool := q.AnswerBool(xInt)
				answer := append([]bool{false}, xBool...)
				answerCh <- answer
			}
		}()
	}

	return answerCh
}

func (s *Srg) IsSolved() bool{
	return s.toFill == 0
}