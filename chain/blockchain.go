package chain


/**
 *	定义区块链这个结构体，用于存储产生的区块（内存中）
 */
type BlockChain struct {
	Blocks []Block
}

/**
 *	创建一个区块链实例，该实例携带一个创世区块
 */
func CreateChainWithGenesis(genesisData []byte) BlockChain {
	genesis := CreateGenesisBlock(genesisData)
	blocks := make([]Block,0)
	blocks = append(blocks,genesis)
	return BlockChain{blocks}
}

func (chain *BlockChain) AdddNewBlock(data []byte)  {
	//1、找到最后一个区块（最新区块）
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	//2、根据最后一个区块产生一个新的区块
	newBlock := CreateBlock(lastBlock.Height,lastBlock.Hash,data)
	//3、把最新产生的区块放入到切片当中
	chain.Blocks = append(chain.Blocks, newBlock)
}