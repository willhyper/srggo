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

func DotSrc(src []bool, indices1, indices2[]int) int{
	N := len(indices1)
	sum:=0
	for i:=0;i<N;i++{
		tf := src[indices1[i]] && src[indices2[i]]
		if tf{
			sum+=1
		}
	}
	return sum
}


func Sum(src []bool) int{
	sum:=0
	for _, v := range src{
		if v{
			sum+=1
		}
	}
	return sum
}