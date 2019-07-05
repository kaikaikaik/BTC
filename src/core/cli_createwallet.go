package core

import "fmt"

//创建钱包账户
func (cli *CLI) createwallet() {
	wallets, _ := NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}
