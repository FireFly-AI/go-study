package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64

	wg sync.WaitGroup

	mutex sync.Mutex
)
func incCounter(i int){
	defer wg.Done()

	for count:= 0; count< i; count++  {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
func main(){
	wg.Add(2)

	go incCounter(2)
	go incCounter(4)

	wg.Wait()

	fmt.Printf("counter fannal is : %d", counter)
}