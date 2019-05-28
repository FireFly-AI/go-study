package main

import "fmt"

const (
    i=1<<iota
    j=3<<iota
    k
    l
)
func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            ii          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,ii)
	
	fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}