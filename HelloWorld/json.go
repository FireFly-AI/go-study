// json解析之 marshal  marshalindent
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main(){

	// 创建一个保存键值对的映射
	c := make(map[string] interface{})
	c["name"] = "xq"
	c["age"] = "22"
	c["contact"] = map[string]interface{}{
		"home":"chengdu",
		"email":"777.qq,com",
	}
	// 映射转换成json串
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	fmt.Println(string(data))

}