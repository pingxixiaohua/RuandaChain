package consensus

import (
	"RuandaChain/chain"
	"RuandaChain/consensus/pos"
	"RuandaChain/consensus/pow"
	"math/big"
)

/**
 *	共识机制的接口标准，用于定义共识方案的接口
 */
type Consensus interface {
	SearchNonce() ([32]byte,int64)

}

func NewProofWork(block chain.Block) Consensus {
	init := big.NewInt(1)
	init.Lsh(init,255-pow.DIFFICULTY)
	return pow.ProofWork{block,init}
}

func NewProofStock() Consensus {
	return pos.ProofStock{}
}
