package main

import (
	"RuandaChain/chain"
	"fmt"
	"github.com/bolt-master"
)

const DBFILE  = "ruanda.db"

/**
 *	项目的主入口
 */
func main() {

	fmt.Println("hello world")

	engine, err := bolt.Open(DBFILE, 0600,nil)
	if err != nil {
		panic(err.Error())
	}

	blockChain := chain.NewBlockChain(engine)
	//创世区块
	blockChain.CreateGenesis([]byte("hello world"))
	//新增一个区块
	err = blockChain.AdddNewBlock([]byte("hello"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//获取最新区块
	//lastBlock, err := blockChain.GetLastBlock()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(lastBlock)

	allBlocks, err := blockChain.GetAllBlocks()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, block := range allBlocks{
		fmt.Println(block)
	}
}

