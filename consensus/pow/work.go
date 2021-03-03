package pow

import (
	"RuandaChain/chain"
	"RuandaChain/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 256位二进制
//给一个大整数，初始值为1，根据自己需要的难度进行左移，左移的位数是256-0的个数
const DIFFICULTY = 10 //初始难度为10，即大整数的开头有10个零

/**
 *	工作量证明
 */
type ProofWork struct {
	Block chain.Block
	Target *big.Int
}


/**
 *	实习共识机制接口的方法
 */
func (work ProofWork) SearchNonce() ([32]byte,int64) {
	fmt.Println("这里是PoW的方法的代码实现过程")
	//block -> nonce
	//block哈希 小于 系统提供的某个目标值

	//1、给定一个nonce值，计算带有nonce的区块哈希
	var nonce int64
	nonce = 0
	for {
		hash := CalculateBlockHash(work.Block, nonce)

		//2、系统给定的值
		target := work.Target

		//3、那步骤1和步骤2比较
		result := bytes.Compare(hash[:],target.Bytes())

		//4、判断结果，区块哈希 < 给定值，返回nonce值
		if result == -1 {
			return hash,nonce
		}//若不满足，则改变nonce值继续试
		nonce++
	}

}


/**
 *	根据当前的区块和当前的nonce值，计算区块的hash值
 */
func CalculateBlockHash(block chain.Block, nonce int64) [32]byte {
	heightByte, _ := utils.Int2Byte(block.Height)
	versionByte, _ := utils.Int2Byte(block.Versionn)
	timeByte, _ := utils.Int2Byte(block.Timestamp)
	nonceByte, _ := utils.Int2Byte(nonce)
	bk := bytes.Join([][]byte{heightByte,versionByte,block.PreHash[:],timeByte,nonceByte,block.Data}, []byte{})

	return sha256.Sum256(bk)

}