package linalg

import (
	"fmt"
)

func betterSolver(q *Question, ch chan []int){
	defer close(ch)
	
	fmt.Println(q)
	fmt.Println("!!!!!")

	qcc := q.CompressCols()
	fmt.Println(qcc)


	ch <- make([]int,3)

}
