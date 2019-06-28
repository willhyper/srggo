package main

// 10,6,3,4
// [[0, 1, 1, 1, 1, 1, 1, 0, 0, 0],
//  [1, 0, 1, 1, 1, 0, 0, 1, 1, 0],
//  [1, 1, 0, 0, 0, 1, 1, 1, 1, 0],
//  [1, 1, 0, 0, 1, 1, 0, 1, 0, 1],
//  [1, 1, 0, 1, 0, 0, 1, 0, 1, 1],
//  [1, 0, 1, 1, 0, 0, 1, 1, 0, 1],
//  [1, 0, 1, 0, 1, 1, 0, 0, 1, 1],
//  [0, 1, 1, 1, 0, 1, 0, 0, 1, 1],
//  [0, 1, 1, 0, 1, 0, 1, 1, 0, 1],
//  [0, 0, 0, 1, 1, 1, 1, 1, 1, 0]]

import (
	"srg"
	"fmt"
	"linalg"
)

func main() {
	s := srg.NewFilledSrg(10, 6, 3, 4)
	s.Fill([]bool{false,false,false,true,true,true,true,false})
	s.Println()

	fmt.Println(" ", s.PivotQuota())
	fmt.Println("-", s.PivotProduct())
	fmt.Println(" ", s.PivotQuotaLeft())

	A, b, quota := s.Question()

	for answer := range linalg.Solver(A, b, quota){
		fmt.Println(answer)
	}

}
