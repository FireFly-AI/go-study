package main

import "fmt"
// 消费者
func consumer(data chan int, done chan bool) {
	for x := range data {
		fmt.Println("recv data : %d", x)
	}
	done <- true
}

func producer(data chan int){
	for i := 0; i < 5; i++ {
		data <- i
	}
	close(data)
}
func main(){
	done := make(chan bool)
	data := make(chan int)

	go consumer(data, done)
	go producer(data)

	<-done
}