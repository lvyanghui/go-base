package chapter05

import (
	"golang.org/x/net/html"
	"fmt"
	"os"
	"net/http"
)

type Node struct {
	Type         NodeType
	Data         string
	Attr         []Attribute
	FirstChild,NextSibling *Node
}

type NodeType int32

const(
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}


func visit(links []string, n *html.Node) []string{
	if n.Type == html.ElementNode && n.Data == "a"{
		for _,a := range n.Attr{
			if a.Key == "href"{
				links = append(links, a.Val)
			}
		}
	}
	for c:= n.FirstChild; c!=nil; c= c.NextSibling{
		links = visit(links,c)
	}
	return links
}

func outline(stack []string, n *html.Node){
	if n.Type == html.ElementNode{
		stack = append(stack,n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil ; c = c.NextSibling{
		outline(stack,c)
	}
}

func PrintHtml(){
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil{
		fmt.Fprintf(os.Stderr, "fetch:%v\n",err)
		os.Exit(1)
	}
	fmt.Printf("%s",resp.Status)
	doc,err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil{
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	}
	for _, link := range visit(nil,doc){
		fmt.Println(link)
	}

	outline(nil, doc)
}
