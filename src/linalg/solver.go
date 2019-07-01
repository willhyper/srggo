package linalg

import (
	"array"
	"itertools"
)

func Solver(q *Question) (<-chan []int){
	m := q.A
	b := q.b

	ch := make(chan []int)

	go func(){
		defer close(ch)

		pool:= make([]int, m.C)
		for i:=0;i<m.C;i++{
			pool[i]=i
		}

		//quota := m.C // this can be improved if knowing the b value of the row of all one
		quota := b[m.R-1] // todo: remove this convention of appending the condition at the last element of vector b
		x := make([]int, m.C)
		for indices := range itertools.Combinations(pool, quota){
			
			for i:=0 ; i < m.C ; i++ { x[i] = 0} // clear x
			for _, v:= range indices { x[v] = 1 } // fill true from indices

			r := 0
			for ; r < m.R ; r++ {
				prod := array.Dot(m.Row(r), x)
				if b[r] != prod {break}
			}
			
			if r == m.R {				
				answer := append([]int(nil), x...)
				ch <- answer
			}

		}
	}()

	return ch
}