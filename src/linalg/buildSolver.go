package linalg

import (
	"array"
)

func BuildSolver(numOfWorker int) func(*Question) chan *Answer {
	QChan := make(chan *Question, 10000)
	// closure keeps QChann always open
	// buffer 10000 to be investigated
	return func(q *Question) chan *Answer {
		AnsChan := make(chan *Answer)
		// AnsChan is created and closed in every invocation
		QChan <- q
		go func() {
			defer close(AnsChan)
			for {
				if len(QChan) > 0 {
					q := <-QChan

					// there is still unknown => simplify and fork
					// there is no unknown but b!=0, not solvable, discard
					// there is no unknow and b ==0, this is answer so put into AnsChan
					Simplify(q)
						
					if q.IsSolved() {
						AnsChan <- q.x
					} else if q.x.Unknown() {
						if Solvable(q) {
							Fork(q, QChan)
						}
					}
				} else {
					break
				}
			}
		}()
		return AnsChan
	}
}

// note this has been under assumption of there is unknown.
func Solvable(q *Question) bool {
	_true := array.AllNonnegative(q.b)
	if !_true {return false}

	_true = q.A.C > 0 // matrix is not empty
	if !_true {return false}

	_true = array.AllNonnegative(q.upperBound)
	if !_true {return false}

	return true
}