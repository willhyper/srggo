package linalg

import (
	"array"
	"fmt"
)

const UNKNOWN = -1

type Answer struct {
	arr      []int
	location []int
	length   int
}

func NewAnswer(arr, location []int, length int) *Answer {
	a := Answer{
		arr:      arr,
		location: location,
		length:   length,
	}
	a.assert(location[0] == 0, "location[0]!=0")
	N := len(arr)
	a.assert(N > 0, "len(arr)=0")
	a.assert(len(location) == N, "len(location)!=len(arr)")
	locationLast := location[N-1]
	a.assert(locationLast < length, "locationLast > len")
	a.assert(isSorted(location), "location is not sorted")
	return &a
}

func NewAnswerDefault(length int) *Answer {
	arr := make([]int, length)
	location := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = UNKNOWN
		location[i] = i
	}
	return NewAnswer(arr, location, length)
}

func (a *Answer) assert(condition bool, errorMessage string) {
	if !condition {
		fmt.Println(a)
		panic(errorMessage)
	}
}

func isSorted(arr []int) bool {
	N := len(arr)
	for i := 1; i < N; i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func (a *Answer) String() string {
	return fmt.Sprintf("%v\n%v %v\n", a.arr, a.location, a.length)
}

func (a *Answer) Binarize() []bool {

	a.assert(!a.Unknown(), "not all values are known")

	ans := make([]bool, a.length)
	p := 0
	for i := 0; i < len(a.arr); i++ {
		numOfTrues := a.arr[i]
		numOfCols := a.location[i]
		for j := p; j < p+numOfTrues; j++ {
			ans[j] = true
		}
		p = numOfCols
	}
	return ans
}

func (a *Answer) quota() []int {
	locations := append(a.location, a.length)
	return array.Diff(locations)
}

func (a *Answer) Unknown() bool {
	for _, v := range a.arr {
		if v == UNKNOWN {
			return true
		}
	}
	return false
}
