package linalg

import (
	"array"
	"itertools"
	"mat"
)

func Solver(A []bool, b []int, quota int) (<-chan []bool){
	m := mat.NewMatrix(A, len(b))

	ch := make(chan []bool)

	go func(){
		defer close(ch)

		pool:= make([]int, m.C)
		for i:=0;i<m.C;i++{
			pool[i]=i
		}


		x := make([]bool, m.C)
		for indices := range itertools.Combinations(pool, quota){
			
			for i:=0 ; i < m.C ; i++ { x[i] = false} // clear x
			for _, v:= range indices { x[v] = true } // fill true from indices

			r := 0
			for ; r < m.R ; r++ {
				prod := array.Dot(m.Row(r), x)
				if b[r] != prod {break}
			}
			
			if r == m.R {				
				answer := append([]bool(nil), x...)
				ch <- answer
			}

		}
	}()

	return ch
}