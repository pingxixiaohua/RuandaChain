package main

import (
	"RuandaChain/chain"
	"RuandaChain/client"
	"github.com/bolt-master"
)

const DBFILE = "ruanda.db"

/**
 *	项目的主入口
 */
func main() {

	engine, err := bolt.Open(DBFILE, 0600,nil)
	if err != nil {
		panic(err.Error())
	}
	blockChain := chain.NewBlockChain(engine)

	cli := client.Client{blockChain}
	//cli.Help()
	cli.Run()
}

