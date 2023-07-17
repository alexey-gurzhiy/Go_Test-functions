package main

import (
	"EthJar/app/log"
	"EthJar/node/connect"
	"flag"
)

const (
	transactionID  string = "0x3cf25c6a0a4333d33276b1eee17ef7495fc8a88cdf86123c890af03d0f867ca3"
	nodeURL        string = "https://mainnet.infura.io/v3/bfdf85f781ae493aaf2773a6b495c36a"
	nodeURLGoerli  string = "https://goerli.infura.io/v3/bfdf85f781ae493aaf2773a6b495c36a"
	nodeURLSepolia string = "https://sepolia.infura.io/v3/bfdf85f781ae493aaf2773a6b495c36a"
)

func main() {
	var loglevel string

	flag.StringVar(&loglevel, "log level", "Trace", "Specify desired log level")
	flag.Parse()

	logger := log.NewLogger(loglevel)
	app := connect.NewApp(logger)
	app.LiveConnectionToEthereumNode(nodeURLSepolia)
}
