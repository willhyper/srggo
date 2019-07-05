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
	"srg"
)

func main() {
	s := srg.NewFilledSrg(10, 6, 3, 4)

	srgPool := make(chan *srg.Srg, 1000)
	srgPool <- s

	s = <-srgPool
	for possibleRow := range srg.Solve(s) {
		newSrg := s.Copy()
		newSrg.Fill(possibleRow)
		srgPool <- newSrg
	}
	// gopher(srgPool)
}

func gopher(srgPool chan *srg.Srg) {

	for {
		select {
		case s := <-srgPool:
			if s.IsSolved() {
				fmt.Println(s)
			} else {
				for possibleRow := range srg.Solve(s) {
					newSrg := s.Copy()
					newSrg.Fill(possibleRow)
					srgPool <- newSrg
				}
			}
		default:
			return
		}
	}
}
