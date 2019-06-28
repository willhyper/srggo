package linalg

import (
	"array"
	"itertools"
)

func Solver(A []bool, b []int, quota int) (<-chan []bool){
	RC := len(A)
	R:=len(b)
	C:=RC/R

	ch := make(chan []bool)

	go func(){
		defer close(ch)

		pool:= make([]int, C)
		for i:=0;i<C;i++{
			pool[i]=i
		}


		candidate := make([]bool, C)
		for indices := range itertools.Combinations(pool, quota){
			// clear candidate
			for i:=0;i<C;i++{candidate[i]=false}

			// fill true from indices
			for _, i:= range indices{candidate[i]=true}

			//candidate ready. test it with b
			row := make([]bool, C)
			rind :=0
			for i:= 0; i<RC;i+=C{
				row = A[i:i+C]
				prod := array.Dot(row, candidate)
				if b[rind]!=prod {break}
				rind++
			}

			if rind ==R{
				answer := append([]bool(nil), candidate...)
				ch <- answer
			}

		}
	}()

	return ch
}