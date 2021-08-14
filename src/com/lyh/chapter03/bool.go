package chapter03

import "fmt"

func BoolPrint(b bool) int{
	c := bool(false)
	fmt.Println(c)
	if b{
		return 1
	}
	return 0
}
