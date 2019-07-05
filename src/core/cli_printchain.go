package core

import (
	"fmt"
	"strconv"
)

//打印区块
func (cli *CLI) printChain() {
	bc := NewBlockChain("")
	defer bc.Db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("========== Block %x =========\n", block.Hash)
		fmt.Printf("PrevHash:%x\n", block.PrevBlockHash)

		fmt.Printf("Hash    :%x\n", block.Hash)

		pow := NewProofOfWork(block)
		fmt.Printf("POW:%s\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
