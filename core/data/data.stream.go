package data

import (
	// "io/ioutil"
	// "net/http"
	// "time"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xtordoir/goanda/api"
	"github.com/xtordoir/goanda/models"
)

// demo: 174528b8942f833285a1081281b6a2bd-3e7e1c8773651ed8c11cece3e4be100b
//101-001-16314125-001

const (
	API_URL = "https://api-fxtrade.oanda.com"
	//https://api-fxpractice.oanda.com/v3/instruments/EUR_USD/candles?count=6&price=M&granularity=S5
	PRACTICE_STREAM_API = "https://stream-fxpractice.oanda.com"
	TOKEN               = "0ca0727c1936d70205477d488cdc3945-4833533d524c983e3c2195763b0f5dff"
	ACCOUNT_ID          = "001-001-10063432-001"
	APP_NAME            = "MVRQ-MACOS"
)

var closeStream = make(chan bool)

var ctx = api.Context{
	ApiURL:       API_URL,
	StreamApiURL: PRACTICE_STREAM_API,
	Token:        TOKEN,
	Account:      ACCOUNT_ID,
	Application:  APP_NAME,
}

var apiClient = ctx.CreateAPI()

var streamClient = ctx.CreateStreamAPI()

var priceChan = make(chan models.ClientPrice)
var heartbeatChan = make(chan models.PricingHeartbeat)

// func Watch() {
// 	fmt.Println("Watching")
// 	go streamClient.PricingStream([]string{"SPX500_USD"}, priceChan, heartbeatChan)

// 	for cp := range priceChan {
// 		fmt.Println(cp)
// 	}
// 	ticker := time.NewTicker(5 * time.Second)
// 	defer ticker.Stop()
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				// Call your function here
// 				// queryCandles()

// 			case <-closeStream:
// 				ticker.Stop()
// 				fmt.Println("\n Stream closed")
// 			}
// 		}
// 	}()
// }

// func QuitWatching() {
// 	closeStream <- true
// }

// func QueryCandleHistory() []models.CandleStick {
	func QueryCandleHistory() []FrontEndCandle {
	jsonFile, err := os.Open("/Users/octavian/Desktop/sandbox/desktop.rythm/core/data/mock.data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened mock.data.json")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var candles []FrontEndCandle

	json.Unmarshal(byteValue, &candles)

	// res, err := apiClient.GetCandles("EUR_USD", 365, "D")

	// if err != nil {
	// 	panic(err)
	// }

	// return res.Candles
	return candles
}

func QueryAskPrice() {
	res, err := apiClient.GetPricing([]string{"EUR_USD"})

	if err != nil {
		panic(err)
	}

	lastAsk := res.Prices[len(res.Prices)-1].Asks[len(res.Prices[len(res.Prices)-1].Asks)-1].Price

	fmt.Println(lastAsk)
}
