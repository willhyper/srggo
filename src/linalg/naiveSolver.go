package linalg

import (
	"array"
	"itertools"
	"fmt"
)

func naiveSolver(q *Question, ch chan *Answer){
	defer close(ch)

	m := q.A
	b := q.b

	pool:= make([]int, m.C)
	for i:=0;i<m.C;i++{
		pool[i]=i
	}
	location := append([]int(nil), pool...)

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
			answerInt := append([]int(nil), x...)
			answer := NewAnswer(answerInt, location, m.C) //broken since using Answer struct. todo
			fmt.Println(answer)
			ch <- answer
		}

	}
}