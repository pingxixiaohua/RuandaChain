package pow

import "fmt"

type PoW struct {

}

func (pow PoW) Run() interface{} {
	fmt.Println("这是PoW方式的共识实现")
	return nil
}
