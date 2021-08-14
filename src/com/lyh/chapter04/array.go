package chapter04

import (
	"fmt"
	"crypto/sha256"
)

func PrintArray(){

	var h [3]int
	fmt.Println(h[0])
	fmt.Println(h[len(h) -1 ])

	for i,v := range h{
		fmt.Printf("%d %d\n", i , v)
	}

	for _,v := range h{
		fmt.Printf("%d\n" , v)
	}

	var q [3]int = [3]int{1,2}
	fmt.Println(q[2])
	fmt.Printf("%T\n",q)

	p := [3]int{1,2,3}
	//p = [4]int{1,2,3,4}
	fmt.Println(p)

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR : "€", GBP: "￡", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
	r := [...]int{99, -1}
	fmt.Println(len(r))

	a := [2]int{1,2}
	b := [...]int{1,2}
	c := [2]int{1,3}
	fmt.Println(a == b, a == c, b == c)
	//d := [3]int{1,3}
	//fmt.Println(a == d)

	e := []byte("x")
	fmt.Println(e)
	C1:=sha256.Sum256([]byte("x"))
	C2:=sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n",C1, C2, C1 == C2, C1)
}

func zero(ptr *[32]byte){
	*ptr = [32]byte{}
}