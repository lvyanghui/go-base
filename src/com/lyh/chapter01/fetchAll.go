package chapter01

import (
	"time"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func fetch(url string, ch chan <- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil{
		ch <- fmt.Sprintf("while reading %s: %v\n", url,err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}
//http://www.baidu.com/ http://www.sina.com http://www.taobao.com
func HttpAll(){
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:]{
		go fetch(url, ch)
	}
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
