package main

import "fmt"
type user struct {
	name string
	age int
}

type manager struct {
	user					// 匿名嵌入
	title string
}

// 方法

func (u user) Tostring() string{
	return fmt.Sprintf("%+v", u)
}

func main(){
	var mg manager
	mg.name = "Tom"
	mg.age = 18
	mg.title = "CTO"
	fmt.Println(mg);

	// 调用匿名字段的方法 相当于类继承
	fmt.Println(mg.Tostring())
}