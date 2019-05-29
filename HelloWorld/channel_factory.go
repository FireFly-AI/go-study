package main
/*这里采用工厂模式来生产receiver对象*/
import (
	"fmt"
	"sync"
)
type receiver struct {
	sync.WaitGroup
	data chan int
}
// 工厂
func newReceiver() *receiver  {
	r := &receiver{
		data:make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x:= range r.data {
			fmt.Println("recv: ", x)
		}
	}()
	return r
}

func main(){
	r:= newReceiver()
	r.data <- 2
	r.data <- 5
	close(r.data)
	r.Wait()
}
