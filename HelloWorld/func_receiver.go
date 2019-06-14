// 方法（包含接受者的函数叫做方法） 接收者分为：值接收者和指针接收者
package main

import "fmt"

// 定义user类型
type user struct {
	name string
	email string
}

// 值接收者
func (u user) notify() {
	fmt.Printf("name: %s  email: %s \n", u.name, u.email);
}

//  指针接收者
func (u *user)changeEmail(e string) {
	u.email = e;
}

func main(){
	bill := user{name:"bill", email:"555@qq.com"};
	bill.notify();

	lisa := &user{name:"lisa", email:"888@qq.com"};
	lisa.notify();

	bill.changeEmail("666@126.com");
	bill.notify();

	lisa.changeEmail("777@126.com");
	lisa.notify();
}

