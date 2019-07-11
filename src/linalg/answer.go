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
	location := make([]int, length)
	for i := 0; i < length; i++ {
		location[i] = i
	}
	arr := array.Ones(length, UNKNOWN)
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
	for i, p := range a.location {
		numOfTrue := a.arr[i]
		for i:=0;i<numOfTrue;i++{
			ans[p+i]=true
		}
	}
	return ans
}


func (a *Answer) Unknown() bool {
	for _, v := range a.arr {
		if v == UNKNOWN {
			return true
		}
	}
	return false
}

func (a *Answer) UnknownIndicies() []int {
	ind := make([]int, len(a.arr))
	p:=0
	for i, v := range a.arr {
		if v == UNKNOWN {
			ind[p] = i
			p++
		}
	}
	return ind[:p]
}


func (a *Answer) KnownIndicies() []int {
	ind := make([]int, len(a.arr))
	p:=0
	for i, v := range a.arr {
		if v != UNKNOWN {
			ind[p] = i
			p++
		}
	}
	return ind[:p]
}

func (a *Answer) Copy() (*Answer) {
	return &Answer{
		arr:      append([]int(nil), a.arr...),
		location: append([]int(nil), a.location...),
		length:   a.length,
	}
}

func (a *Answer) CopySelect(indices []int) *Answer {
	arr := make([]int, len(indices))
	loc := make([]int, len(indices))

	for inew, ioriginal := range indices {
		arr[inew] = a.arr[ioriginal]
		loc[inew] = a.location[ioriginal]
	}
	return &Answer{
		arr:      arr,
		location: loc,
		length:   a.length,
	}
}


func (a *Answer) FillUnknown(ans, nth int) {
	unknownIndexToFill := array.FindIntIndex(a.arr, UNKNOWN, nth)
	a.arr[unknownIndexToFill] = ans
}

