package martingalestrategy

import (
	"testing"

	"github.com/fedealliani/gobot/bot"
	"github.com/stretchr/testify/assert"
)

var martingaleStrategyDefault = MartingaleStrategy{
	AmountFirstTrade: 100,
}

var martingaleStrategyAfterFirstBuy = MartingaleStrategy{
	AmountFirstTrade: 100,
	BuyMode:          true,
	Init:             true,
	Purchases:        []float64{100},
	Averages:         []float64{100},
	Bets:             []float64{100},
	Iteration:        0,
}

func TestShouldBuy01WhenInitBuy(t *testing.T) {
	info := bot.Info{AmountQuoteAssetCoin: martingaleStrategyDefault.AmountFirstTrade * 2, AvgPrice: 10}
	res, amount, err := martingaleStrategyDefault.ShouldBuy(info)
	if !assert.Equal(t, nil, err) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, true, res) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, martingaleStrategyDefault.AmountFirstTrade, amount) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
}

func TestShouldBuy02WhenDontHaveMoneyDontBuy(t *testing.T) {
	info := bot.Info{AmountQuoteAssetCoin: 0, AvgPrice: 10}
	res, amount, err := martingaleStrategyDefault.ShouldBuy(info)
	if !assert.Equal(t, nil, err) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, false, res) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, float64(0), amount) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
}

func TestShouldBuy03WhenLostBuyDouble(t *testing.T) {
	info := bot.Info{AmountQuoteAssetCoin: 5 * martingaleStrategyAfterFirstBuy.AmountFirstTrade, AvgPrice: 10}
	res, amount, err := martingaleStrategyAfterFirstBuy.ShouldBuy(info)
	if !assert.Equal(t, nil, err) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, true, res) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, 2*martingaleStrategyAfterFirstBuy.AmountFirstTrade, amount) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
}

func TestShouldBuy04WhenWinDontBuy(t *testing.T) {
	info := bot.Info{AmountQuoteAssetCoin: 5 * martingaleStrategyAfterFirstBuy.AmountFirstTrade, AvgPrice: 10000}
	res, amount, err := martingaleStrategyAfterFirstBuy.ShouldBuy(info)
	if !assert.Equal(t, nil, err) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, false, res) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, float64(0), amount) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
}

func TestShouldBuy05WhenAddComissionAndDontHaveMoneyDontBuy(t *testing.T) {
	info := bot.Info{AmountQuoteAssetCoin: martingaleStrategyAfterFirstBuy.AmountFirstTrade + 1, AvgPrice: martingaleStrategyAfterFirstBuy.AmountFirstTrade, Comission: 0.075 / 100}
	res, amount, err := martingaleStrategyDefault.ShouldBuy(info)
	if !assert.Equal(t, nil, err) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, false, res) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, float64(0), amount) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
}
