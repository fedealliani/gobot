package percentagestrategy

import (
	"testing"

	"github.com/fedealliani/gobot/bot"
	"github.com/stretchr/testify/assert"
)

var percentageStrategyDefault = PercentageStrategy{
	Percentage: 1,
}

var percentageStrategyAfterInit = PercentageStrategy{
	Percentage: 1,
	Init:       true,
	BuyMode:    true,
	Target:     5000,
}

var percentageStrategyAfterFirstBuy = PercentageStrategy{
	Percentage: 1,
	Init:       true,
	BuyMode:    false,
	Target:     5000,
}

func TestShouldBuy01WhenInitDontBuy(t *testing.T) {
	info := bot.Info{AmountQuoteAssetCoin: 10000, LastPrice: 500}
	res, amount, err := percentageStrategyDefault.ShouldBuy(info)
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

func TestShouldBuy02DontBuyBecauseDontMeetExpectedPercentage(t *testing.T) {
	lastPrice := percentageStrategyAfterInit.Target
	info := bot.Info{AmountQuoteAssetCoin: 10000, LastPrice: lastPrice}
	res, amount, err := percentageStrategyAfterInit.ShouldBuy(info)
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

func TestShouldBuy03BuyBecauseMeetExpectedPercentage(t *testing.T) {
	lastPrice := percentageStrategyAfterInit.Target - 1
	info := bot.Info{AmountQuoteAssetCoin: 10000, LastPrice: lastPrice}
	res, amount, err := percentageStrategyAfterInit.ShouldBuy(info)
	if !assert.Equal(t, nil, err) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, true, res) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
	if !assert.Equal(t, float64(info.AmountQuoteAssetCoin), amount) {
		t.Fatalf("There was an error on ShouldBuy %v", err)
	}
}

func TestShouldSell01WhenInitDontSell(t *testing.T) {
	info := bot.Info{AmountQuoteAssetCoin: 10000, LastPrice: 500}
	res, amount, err := percentageStrategyDefault.ShouldSell(info)
	if !assert.Equal(t, nil, err) {
		t.Fatalf("There was an error on ShouldSell %v", err)
	}
	if !assert.Equal(t, false, res) {
		t.Fatalf("There was an error on ShouldSell %v", err)
	}
	if !assert.Equal(t, float64(0), amount) {
		t.Fatalf("There was an error on ShouldSell %v", err)
	}
}

func TestShouldSell02DontSellBecauseDontMeetExpectedPercentage(t *testing.T) {
	lastPrice := percentageStrategyAfterFirstBuy.Target
	info := bot.Info{AmountQuoteAssetCoin: 10000, LastPrice: lastPrice}
	res, amount, err := percentageStrategyAfterFirstBuy.ShouldSell(info)
	if !assert.Equal(t, nil, err) {
		t.Fatalf("There was an error on ShouldSell %v", err)
	}
	if !assert.Equal(t, false, res) {
		t.Fatalf("There was an error on ShouldSell %v", err)
	}
	if !assert.Equal(t, float64(0), amount) {
		t.Fatalf("There was an error on ShouldSell %v", err)
	}
}

func TestShouldSell03SellBecauseDontMeetExpectedPercentage(t *testing.T) {
	lastPrice := percentageStrategyAfterFirstBuy.Target + 1
	info := bot.Info{AmountQuoteAssetCoin: 10000, LastPrice: lastPrice}
	res, amount, err := percentageStrategyAfterFirstBuy.ShouldSell(info)
	if !assert.Equal(t, nil, err) {
		t.Fatalf("There was an error on ShouldSell %v", err)
	}
	if !assert.Equal(t, true, res) {
		t.Fatalf("There was an error on ShouldSell %v", err)
	}
	if !assert.Equal(t, float64(info.AmountBaseAssetCoin), amount) {
		t.Fatalf("There was an error on ShouldSell %v", err)
	}
}
