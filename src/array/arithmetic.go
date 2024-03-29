package array


func DotSrc(src []bool, indices1, indices2[]int) int{
	N := len(indices1)
	if len(indices2)!=N{
		panic("len(indices1)!=len(indices2)")
	}

	sum:=0
	for i:=0;i<N;i++{
		tf := src[indices1[i]] && src[indices2[i]]
		if tf{
			sum+=1
		}
	}
	return sum
}

func SumBool(src []bool) int{
	sum:=0
	for _, v := range src{
		if v{
			sum+=1
		}
	}
	return sum
}

func Diff(a []int) []int{

	if len(a)<=1{
		panic("len(a)<=1. Cannot diff")
	}

	diff := make([]int, len(a)-1)
	for i:=0;i<len(diff);i++{
		diff[i] = a[i+1]-a[i]
	}
	return diff
}


func SubstractVector(a, c []int) []int{
	b := make([]int, len(a))
	for i:=0 ; i< len(a); i++ {
		b[i] = a[i] - c[i]
	}
	return b
}

func Multiply(a []bool, c int) []int{
	b := make([]int, len(a))
	for i:=0 ; i< len(a); i++ {
		if a[i]{
			b[i] = c
		}
	}
	return b
}