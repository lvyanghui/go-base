package chapter01

import (
	"fmt"
	"os"
	"strings"
)


func OsParam(){
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func OsParam1(){
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func OsParam2(){

	fmt.Println(os.Args[1:])
}

func OsParam3(){

	fmt.Println(strings.Join(os.Args[1:]," "))
}

func Practice1(){

	fmt.Println(os.Args[0])
	os.Args[0] = "Practice"
	fmt.Println(os.Args[0])
}

func Practice2(){

	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("i = %d,param= %s ", i , os.Args[i])
		fmt.Println()
	}
}


/**
暂时不会用time
 */
func Practice3(){


}