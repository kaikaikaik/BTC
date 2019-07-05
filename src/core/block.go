package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

//区块链中的一个区块
type Block struct {
	Timestamp     int64          //时间戳
	Transactions  []*Transaction //交易数据
	PrevBlockHash []byte         //前一个区块的哈希值
	Hash          []byte         //区块自身的哈希值，用于校验区块数据有效
	Nonce         int            //Nonce随机数，用于工作量的证明
}

//创建一个新的区块 传入交易数据和前一个哈希 返回一个新的区块
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		transactions,
		prevBlockHash,
		[]byte{},
		0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//创建创世块 返回一个创世块
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

//进行哈希计算 返回区块中交易的哈希值
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]

}

//序列化区块，将区块转为字节数组
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	//编码为字节数组
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//Deserialize deserialize a block
//反序列化区块
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
