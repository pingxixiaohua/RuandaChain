package pos

import "fmt"

type PoS struct {

}

func (pos PoS) Run() interface{} {
	fmt.Println("这里是PoS的共识机制实现方式")
	return nil
}