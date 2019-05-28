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

// 接口定义

type Printer interface {
	Print()
}

// 实现接口
func (u user)Print() {
	fmt.Printf("%+v\n", u)
}

func main(){
	var mg manager
	mg.name = "Tom"
	mg.age = 18
	mg.title = "CTO"
	fmt.Println(mg);

	// 调用匿名字段的接口方法 相当于类继承
	fmt.Println(mg.Tostring())
	mg.Print()

	// 接口使用
	var  u user
	u.name = "XQ"
	u.age = 19
	var p Printer = u
	p.Print()
}