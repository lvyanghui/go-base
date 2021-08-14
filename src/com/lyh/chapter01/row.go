package chapter01

import (
	"bufio"
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

func PrintInput1(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan(){
		counts[input.Text()]++
	}

	for line, n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n",n, line)
		}
	}
}

func PrintInput2(){
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin,counts)
	}else{
		for _, arg := range files{
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n",err)
				continue
			}
			countLines(f,counts)
			f.Close()
		}
	}
	for line, n := range counts{
		if n > 1 {
		fmt.Printf("%d\t%s\n",n, line)
		}
	}
}

func countLines(f * os.File, counts map[string]int){
	input := bufio.NewScanner(f)

	i := 0
	for input.Scan() {
		i = i + 1
		counts[input.Text()] = i
	}
}

//f.txt
func PrintInput3(){
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:]{
		data, err := ioutil.ReadFile(fileName)
		if err != nil{
			fmt.Fprintf(os.Stderr, "dup3: %v\n",err)
			continue
		}
		i := 0
		for _, line := range strings.Split(string(data),"\n"){
			i = i + 1
			counts[line] = i
		}
	}

	for line, n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n",n, line)
		}
	}
}