// io相关输出 stdout
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init(){
	if len(os.Args) != 2 {
		fmt.Printf("Usage : ../example2")
		os.Exit(-1)
	}
}
func main(){
	// 从web服务器得到响应
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// 复制Body到Stdout
	io.Copy(os.Stdout, r.Body)
	if err:= r.Body.Close(); err != nil{
		fmt.Println(err);
	}
}