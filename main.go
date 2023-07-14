package main

import (
	"EthJar/app/log"
	"EthJar/node/connect"
)

const (
	transactionID  string = "0x3cf25c6a0a4333d33276b1eee17ef7495fc8a88cdf86123c890af03d0f867ca3"
	nodeURL        string = "https://mainnet.infura.io/v3/bfdf85f781ae493aaf2773a6b495c36a"
	nodeURLGoerli  string = "https://goerli.infura.io/v3/bfdf85f781ae493aaf2773a6b495c36a"
	nodeURLSepolia string = "https://sepolia.infura.io/v3/bfdf85f781ae493aaf2773a6b495c36a"
)

func main() {
	conn, cntnxt := connect.LiveConnectionToEthereumNode(nodeURLSepolia)
	log.Trace(conn, "\n", cntnxt)

}
