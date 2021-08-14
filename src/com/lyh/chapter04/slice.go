package chapter04

import (
	"fmt"
)

func PrintSlice(){
	months := [...]string{1: " January",2: " February",3: " March",
	4: " April",5: " May",6: " June",
	7: " July",8: " August",9: " September",
	10: " October",11: " November", 12: "December"}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2,summer)

	for _,s := range Q2{
		for _,t := range summer{
			if s == t{
				fmt.Printf("%s appears in both\n",s)
			}
		}
	}
	//fmt.Println(summer[:20])

	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)


	a := [...]int{0,1,2,3,4,5}
	reverse(a[:])
	fmt.Println(a)

	var runes []rune
	for _,r := range "hello,世界"{
		runes = append(runes, r)
	}
	fmt.Printf("%q\n",runes)

	var x,y []int
	for i:= 0; i < 10; i++{
		y = appendInt(x,i)
		fmt.Printf("%d cap=%d\t%v\n",i, cap(y),y)
		x = y
	}

	data := []string{"one","","three"}
	fmt.Printf("%q\n", nonempty2(data))
	//fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	s := []int{5,6,7,8,9}
	fmt.Println(remove(s,3))
	fmt.Println(remove2(s,1))
}

func reverse(s []int){
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1{
		s[i],s[j] = s[j], s[i]
	}
}

func equal(x, y[]string) bool{
	if(len(x) != len(y)){
		return false
	}

	for i := range x{
		if x[i] != y[i]{
			return false
		}
	}
	return true
}

func appendInt(x []int, y int) []int{
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x){
		z = x[:zlen]
	}else{
		zcap := zlen
		if zcap < len(x) * 2{
			zcap = len(x) * 2
		}
		z = make([]int, zlen, zcap)
		copy(z,x)
	}
	z[len(x)] = y
	return z
}

func nonempty(strings []string) []string{
	i := 0
	for _, s := range strings{
		if s != ""{
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}


func nonempty2(strings []string) []string{
	out := strings[:0]
	for _, s := range strings{
		if s != ""{
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int{
	copy(slice[i:len(slice)],slice[i+1:len(slice)])
	return slice[:len(slice) - 1]
}

func remove2(slice []int, i int) []int{
	slice[i] = slice[len(slice) - 1]
	return slice[:len(slice) - 1]
}

