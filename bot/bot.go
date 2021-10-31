package bot

import "github.com/adshao/go-binance/v2"

type GoBot struct {
	strategy      Strategy
	config        Config
	binanceClient *binance.Client
}

func New(strategy Strategy, config Config, binanceClient *binance.Client) *GoBot {
	return &GoBot{
		strategy:      strategy,
		config:        config,
		binanceClient: binanceClient,
	}
}
