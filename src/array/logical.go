package array

func Equal(src []bool, indices1, indices2[]int) bool{
	N:=len(indices1)
	if N!=len(indices2){
		panic("len(indices1)!=len(indices2)")
	}
	for i:=0;i<N;i++{
		if src[indices1[i]] != src[indices2[i]]{
			return false
		} 
	}
	return true
}