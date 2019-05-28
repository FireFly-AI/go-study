package main

import (
	"fmt"
	"time"
)

func testroutine(id int){
	for i:=0;i<5 ;i++  {
		fmt.Printf("id: %d  index: %d \n", id, i)
		time.Sleep(time.Second)
	}
}
func main(){
	go testroutine(1)
	go testroutine(2)

	time.Sleep(time.Second * 6)
}
