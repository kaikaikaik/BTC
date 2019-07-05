package core

import (
	"fmt"
	"log"
)

//创建区块链
func (cli *CLI) createBlockChain(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	bc := CreateBlockChain(address)
	bc.Db.Close()
	fmt.Println("Done!")
}
