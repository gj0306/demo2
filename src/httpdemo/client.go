package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){

	resp,err := http.Get("http://127.0.0.1:9000/")
	if err == nil{
		fmt.Println("结果:",resp)
		if body, err := ioutil.ReadAll(resp.Body);err==nil{
			fmt.Println("内容：",string(body))
		}
	}else {
		fmt.Println("err:%+v",err)
	}
}
