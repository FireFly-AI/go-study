package main

/*这里写了两个chan通道进行传值，done通道检测到关闭后程序退出*/
import (
	"fmt"
)
func main(){
	done := make(chan struct{})
	data := make(chan string)

	go func(){
		s := <-data
		fmt.Println(s)
		close(done)
	}()

	data <- "set data to channel"
	<-done
}

