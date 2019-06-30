package mat

func NewMatrix(arr []bool, R int) *Matrix {
	RC := len(arr)
	C := RC / R

	m := Matrix{
		arr: arr,
		R:   R,
		C:   C,
	}
	return &m
}
