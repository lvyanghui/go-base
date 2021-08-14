package chapter03

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func StringPrint(){
	var s string = "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Println("goodbye" + s[5:])


	s1 := "hello, 世界"
	fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))

	for i := 0; i < len(s1); i++{
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n",i, r)
		i += size
	}

	for i ,r := range s1{
		fmt.Printf("%d\t%q\t%d\n",i, r,r)
	}

	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))

	fmt.Println(basename1("a/b/c.go"))
	fmt.Println(basename1("c.d.go"))
	fmt.Println(basename1("abc"))

}

func basename(s string) string{
	for i := len(s) - 1; i >= 0; i--{
		if s[i] == '/'{
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i--{
		if s[i] == '.'{
			s = s[:i]
			break
		}
	}
	return s
}

//bytes strings strconv unicode
func basename1(s string) string{
	slash := strings.LastIndex(s,"/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, ".");dot >= 0{
		s = s[:dot]
	}
	return s
}
