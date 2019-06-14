// atomic 原子操作用例
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64

	wg sync.WaitGroup
)

func incCounter(i int){
	defer wg.Done()
	for count :=0; count < i; count++ {
		atomic.AddInt64(&counter, 1)
	}
	runtime.Gosched()

}
func main(){
	wg.Add(2)

	go incCounter(10)
	go incCounter(11)

	wg.Wait()

	fmt.Printf("Counter finnal is : %d", counter)
}