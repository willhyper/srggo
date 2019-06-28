package array

import "fmt"

func Sprint(a []bool) string{
	mi := *ToInt(a)
	return fmt.Sprint(mi)
}

func Println(a []bool){
	fmt.Println(Sprint(a))
}