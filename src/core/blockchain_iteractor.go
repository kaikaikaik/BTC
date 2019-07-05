package core

import (
	"github.com/boltdb/bolt"
	"log"
)

//用于在区块链的区块上进行迭代
type BlockChain struct {
	tip []byte
	Db  *bolt.DB
}

//通过迭代器遍历区块信息
func (i *BlockChainIterator) Next() *Block {
	var block *Block

	err := i.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		//反序列化当前的区块
		block = DeserializeBlock(encodedBlock)
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}