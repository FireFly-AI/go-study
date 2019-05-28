package main
/*defer 用来释放资源，解除锁定、或执行一些清理操作*/
import "fmt"

func testdefer(){
	defer fmt.Println("print later...")
	fmt.Println("print...")
}
func main(){
	testdefer();
}
