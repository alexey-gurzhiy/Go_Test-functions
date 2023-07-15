package connect

import (
	"context"

	log "EthJar/app/log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type App struct {
	conn   *ethclient.Client
	logger log.Logger
}

func NewApp(logger log.Logger) *App {
	return &App{
		logger: logger,
	}
}

func (a *App) LiveConnectionToEthereumNode(nodeURL string) (ctx context.Context) {
	var err error
	a.logger.WithFields(logrus.Fields{
		"\n	1. Module":   "node/connect",
		"\n	2. Function": "LiveConnectionToEthereumNode",
	})

	a.conn, err = ethclient.Dial(nodeURL)

	if err != nil {
		a.logger.Fatal("Failed to connect to Eth Node:", err)
	} else {
		a.logger.Trace("connection to ", nodeURL, " was established")
	}
	ctx = context.Background()
	a.logger.Trace(a.conn, "\n", ctx)
	return ctx
}
