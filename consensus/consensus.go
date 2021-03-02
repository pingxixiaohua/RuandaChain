package consensus

import (
	"RuandaChain/consensus/pos"
	"RuandaChain/consensus/pow"
)

/**
 *	Consensus作为接口，用于定义共识机制的标准
 */
type Consensus interface {
	Run() interface{}
}

func NewPoW() Consensus {
	proof := pow.PoW{}
	return proof
}

func NewPoS() Consensus {
	proof := pos.PoS{}
	return proof
}
