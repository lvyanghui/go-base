package chapter01

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

/**
配置program arguments http://www.baidu.com/ http://www.sina.com http://www.taobao.com
 */
//http://www.baidu.com/
func HttpContent(){
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch:%v\n",err)
			os.Exit(1)
		}
		fmt.Printf("%s",resp.Status)
		b,err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:%v\n",url,err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)

	}
}
