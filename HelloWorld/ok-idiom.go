/* ok-idiom（A跌目）模式：多返回值中用一个名为ok的布尔值来标记操作是否成功*/
package main

import "fmt"

func main(){
	m := make(map[string]int)
	m["a"] = 1
	x,ok1 := m["a"]
	y,ok := m["b"]
	fmt.Println("x: %d, ok1: ", x, ok1)
	fmt.Println("y: %d, ok: ", y, ok)
	delete(m, "a")
}
