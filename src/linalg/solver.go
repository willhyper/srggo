package linalg


func Solver(q *Question) (<-chan []int){

	ch := make(chan []int)

	go naiveSolver(q, ch)

	return ch
}