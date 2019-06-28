package array


func GetValues(src []bool, indices []int) *[]bool{
	v := make([]bool, len(indices))
	for i:=0;i<len(indices);i++{
		v[i] = src[indices[i]]
	}
	return &v
}