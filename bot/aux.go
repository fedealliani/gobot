package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func getCandle(cl *binance.Kline) Candle {
	open, err := strconv.ParseFloat(cl.Open, 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	high, err := strconv.ParseFloat(cl.High, 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	low, err := strconv.ParseFloat(cl.Low, 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	close, err := strconv.ParseFloat(cl.Close, 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	volume, err := strconv.ParseFloat(cl.Volume, 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}

	candle := Candle{
		OpenTime:  float64(cl.OpenTime),
		Open:      open,
		High:      high,
		Low:       low,
		Close:     close,
		Volume:    volume,
		CloseTime: float64(cl.CloseTime),
	}
	return candle
}

func GetValueInUSD(symbol string, value float64) (float64, error) {
	result := 0.0

	resp, err := http.Get("https://cex.io/api/last_price/" + symbol + "/USD")
	if err != nil {
		fmt.Printf("there was an error getting candles %v\n", err)
		return result, err
	}
	if resp.StatusCode != 200 {
		fmt.Printf("there was an invalid status code %d", resp.StatusCode)
		return result, err

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	var r UsdValue
	err = json.Unmarshal(body, &r)
	if err != nil { // Parse []byte to the go struct pointer
		fmt.Printf("Cannot unmarshal JSON %v\n", err)
		return result, err
	}

	result, err = strconv.ParseFloat(r.LPrice, 64)
	if err != nil {
		fmt.Printf("Cannot converting usd price to float %v\n", err)
		return result, err
	}
	return result * value, nil
}
func Notify(action string) {
	NotifyTelegram(action)
}

func NotifyTelegram(action string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Create a telegram service. Ignoring error for demo simplicity.
	token := os.Getenv("TELEGRAM_TOKEN")

	telegramService, _ := telegram.New(token)

	// Passing a telegram chat id as receiver for our messages.
	// Basically where should our message be sent?
	chatIDS := os.Getenv("TELEGRAM_CHAT_ID")

	chatID, err := strconv.ParseInt(chatIDS, 10, 64)
	if err != nil {
		panic(err)
	}
	telegramService.AddReceivers(chatID)

	// Create our notifications distributor.
	notifier := notify.New()

	// Tell our notifier to use the telegram service. You can repeat the above process
	// for as many services as you like and just tell the notifier to use them.
	// Inspired by http middlewares used in higher level libraries.
	notifier.UseServices(telegramService)

	// Send a test message.
	_ = notifier.Send(
		context.Background(),
		"Time to making Money",
		action,
	)
}
