package linalg

import (
	"fmt"
)

func betterSolver(q *Question, ch chan *Answer){
	defer close(ch)

	fmt.Println(q)
	qcc := q.CompressCols()
	fmt.Println(qcc)
	// under construction
	
	ch <- NewAnswerDefault(3)

}
