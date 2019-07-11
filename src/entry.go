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
	"fmt"
	"linalg"
	"srg"
)

func main() {

	s := srg.NewFilledSrg(10, 6, 3, 4)
	srgQueue := make(chan *srg.Srg, 1000000) // cannot block sender from adding srg
	srgQueue <- s

	SolveAllTheWay(srgQueue)
	
	fmt.Println("bye")
}

func SolveOneStep(srgQueue chan *srg.Srg) { // for debug

	if len(srgQueue) == 0 {
		fmt.Println("No question in queue")
		return
	}

	solver := linalg.BuildSolver(1)

	s := <- srgQueue
	fmt.Println(s.ToFill())
	if s.ToFill() == 0 { // solved
		fmt.Println(s)
		return
	} else if s.ToFill() == 1 { // solved
		answer := []bool{false}
		s.Fill(answer)
		fmt.Println(s)
		return
	}

	fmt.Println(s)
	q := s.Question() // generate Question form
	fmt.Println(q)

	for ans := range solver(q){
		fmt.Println("ans", ans)

		xBool := ans.Binarize()		
		answer := append([]bool{false}, xBool...)

		newSrg := s.Copy()
		newSrg.Fill(answer)
		srgQueue <- newSrg // no block sender from adding srg
	}

}

func SolveAllTheWay(srgQueue chan *srg.Srg) {

	solver := linalg.BuildSolver(1)

	for {
		if len(srgQueue) > 0 {
			s := <-srgQueue
			if s.ToFill() == 0 { // solved
				fmt.Println(s)
			} else if s.ToFill() == 1 { // solved
				answer := []bool{false}
				s.Fill(answer)
				fmt.Println(s)
			} else {
				q := s.Question() // generate Question form
				for ans := range solver(q) {
					xBool := ans.Binarize()
					answer := append([]bool{false}, xBool...)
					newSrg := s.Copy()
					newSrg.Fill(answer)
					srgQueue <- newSrg // no block sender from adding srg
				}
			}
		} else {
			break
		}
	}
}
