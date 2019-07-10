package array

func GetValues(src []bool, indices []int) *[]bool {
	v := make([]bool, len(indices))
	for i := 0; i < len(indices); i++ {
		v[i] = src[indices[i]]
	}
	return &v
}

func SingleOutElement(a []int, index int) ([]int, int) {
	b := make([]int, len(a)-1)

	for i := 0; i < index; i++ {
		b[i] = a[i]
	}
	v := a[index]
	for i := index + 1; i < len(a); i++ {
		b[i-1] = a[i]
	}
	return b, v
}
