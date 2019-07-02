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
)

func main() {
	s := srg.NewFilledSrg(10, 6, 3, 4)

	questionPool := make(chan *srg.Srg,1000)
	questionPool<-s

	gopher(questionPool)
}

func gopher(questionPool chan *srg.Srg){
	
	for{
		select{
			case q := <-questionPool:
				if q.IsSolved(){
					fmt.Println(q)
				}else{
					for possibleAnswer := range srg.Solve(q){
						newQuestion := q.Copy()
						newQuestion.Fill(possibleAnswer)
						questionPool <- newQuestion
					}
				}
			default:
				return
		}	
	}
}
	

