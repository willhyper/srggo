package array


func Dot(a, b []bool) int{
	sum := 0
	for i:=0;i<len(a);i++{
		tf := a[i] && b[i]
		if tf{
			sum+=1
		}
	}
	return sum
}
