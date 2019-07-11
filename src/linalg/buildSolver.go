package linalg

import (
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

					if q.x.Unknown() {
						Simplify(q)
						Fork(q, QChan)
					} else if q.IsSolved() {
						AnsChan <- q.x
					}
				} else {
					break
				}
			}
		}()
		return AnsChan
	}
}
