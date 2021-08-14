package chapter04

import (
	"fmt"
	"sort"
)

var graph map[string]map[string]bool

func PrintMap(){
	//var ages map[string]int
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	fmt.Println(ages)
	fmt.Println(ages["charlie"])
	delete(ages,"charlie")
	fmt.Println(ages)
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages)

	for name, age := range ages{
		fmt.Printf("%s\t%d\n",name, age)
	}

	var names []string
	for name:= range ages{
		names = append(names,name)
	}
	sort.Strings(names)
	fmt.Println(names)
	/*ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}*/


}

func equalMap(x, y map[string]int) bool{
	if len(x) != len(y){
		return false
	}

	for k,xv := range x{
		if yv, ok := y[k]; !ok || xv != yv{
			return false
		}
	}
	return true
}

func addEdge(from , to string){
	edges := graph[from]
	if edges == nil{
		edges = make(map[string]bool)
		graph[from]= edges
	}
	edges[to]= true
}

func hasEdge(from, to string) bool{
	return graph[from][to]
}
