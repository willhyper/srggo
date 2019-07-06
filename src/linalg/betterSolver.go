package linalg

import (
	"fmt"
)

func betterSolver(q *Question, ch chan *Answer){
	defer close(ch)

	fmt.Println(q)
	q.CompressCols()
	fmt.Println(q)
	// under construction
	

	ch <- NewAnswerDefault(3)

}
