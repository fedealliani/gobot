package bot

type GoBot struct {
	strategy       Strategy
	config         Config
	exchangeClient Exchange
}

func New(strategy Strategy, config Config, exchangeClient Exchange) *GoBot {
	return &GoBot{
		strategy:       strategy,
		config:         config,
		exchangeClient: exchangeClient,
	}
}
