package linalg


func Solve(q *Question) (<-chan *Answer){

	ch := make(chan *Answer)

	//go naiveSolver(q, ch)
	go betterSolver(q, ch)
	
	return ch
}