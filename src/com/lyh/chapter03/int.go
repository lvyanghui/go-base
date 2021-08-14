package chapter03

import "fmt"

func IntPrint(){

	//uint8 uint16 uint32 uint64
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	a := uint64(20)
	fmt.Println(a)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)
}
