package pos

import (
	"RuandaChain/chain"
	"fmt"
)

type ProofStock struct {
	Block chain.Block
}

func (stock ProofStock) SearchNonce() ([32]byte,int64) {
	fmt.Println("这里是PoS机制实现")
	return [32]byte{},0
}
