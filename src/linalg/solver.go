package linalg


func Solve(q *Question) (<-chan []int){

	ch := make(chan []int)

	//go naiveSolver(q, ch)
	go betterSolver(q, ch)
	return ch
}