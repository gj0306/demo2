package main
import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req) //请求体
	fmt.Println(w.Header())
	_,_ = io.WriteString(w, "hello, world!\n") //返回内容
}
func main() {
	http.HandleFunc("/", HelloServer) //分发请求
	for i:=0;i<10;i++{
		err := http.ListenAndServe(":9000", nil) //监听
		if err != nil {
			log.Fatal("ListenAndServe: ", err) //记录错误信息
		}
	}
}
