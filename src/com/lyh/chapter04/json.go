package chapter04

import (
	"encoding/json"
	"log"
	"fmt"
	"time"
	"net/url"
	"strings"
	"net/http"
)

const IssuesURL  = "https://api.github.com/search/issues"
type Movie struct {
	Title string
	Year int  `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issues
}
type Issues struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreateAt time.Time `json:"create_at"`
	Body string
}

type User struct{
	Login string
	HTMLURL string `json:"html_url"`
}

func PrintJson(){
	var movies = []Movie{
		{"cas",1942,false,[]string{"sdfdsf","s0999213"}},
		{"poi",1946,true,[]string{"ttttt","qwqewqe"}},
		{"nba",1923,true,[]string{"jordon","kobe"}},
	}

	//data, err := json.Marshal(movies)

	data, err := json.MarshalIndent(movies,"","    ")
	if err != nil{
		log.Fatalf("JSON marsha1 failed : %s", err)
	}
	fmt.Printf("%s\n",data)

	var titles []struct{Title string}
	if err := json.Unmarshal(data,&titles); err != nil{
		log.Fatalf("JSON Unmarshal failed : %s", err)
	}
	fmt.Println(titles)

	result, err := SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items{
		fmt.Printf("#%-5d %9.9s %.55s\n",item.Number,item.User.Login, item.Title)
	}
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms," "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil{
		return nil,err
	}

	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil, fmt.Errorf("search  query failed: %s",resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil{
		resp.Body.Close()
		return nil ,err
	}
	resp.Body.Close()
	return &result, nil
}
