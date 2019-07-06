package linalg

import (
	"fmt"
)

func Solve(q *Question) (<-chan *Answer){

	ch := make(chan *Answer)

	//go naiveSolver(q, ch)
	go betterSolver(q, ch)
	
	return ch
}


func betterSolver(q *Question, ch chan *Answer){
	defer close(ch)

	fmt.Println(q)
	q.CompressCols()
	fmt.Println(q)
	// under construction
	
	ch <- NewAnswerDefault(3)

}
