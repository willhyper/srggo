package array

func Ones(length int, multiplier int) []int {
	arr := make([]int, length)
	for i:=0;i<length; i++{
		arr[i] = multiplier
	}
	return arr
}