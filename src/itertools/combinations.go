package itertools

func Combinations(iterable []int, r int) (<-chan []int){

	ch := make(chan []int)
	
	go func(){
	
		defer close(ch)
	
		pool := iterable
		n := len(pool)
		indices := make([]int, r)
		result := make([]int, r)

		if r > n {return}

		for i := range indices {indices[i] = i}

		for i, el := range indices {result[i] = pool[el]}

		ch<-append([]int(nil), result...)

		for {
			i := r - 1
			for ; i >= 0 && indices[i] == i+n-r; i -= 1 {}

			if i < 0 {return}

			indices[i] += 1
			for j := i + 1; j < r; j += 1 {indices[j] = indices[j-1] + 1}

			for ; i < r; i += 1 {result[i] = pool[indices[i]]}
			

			ch<-append([]int(nil), result...)

		}
	}() // end of goroutine 

	return ch
}
