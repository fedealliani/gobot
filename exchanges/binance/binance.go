package binance

type BinanceClient struct {
	ApiKey    string
	SecretKey string
}

func New(apiKey string, secretKey string) *BinanceClient {
	return &BinanceClient{ApiKey: apiKey, SecretKey: secretKey}
}
