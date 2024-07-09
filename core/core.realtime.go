package rythm

import (
	// "encoding/json"
	"fmt"
	// "log"
	"time"

	"desktop.rythm/core/data"
	"github.com/xtordoir/goanda/models"
)

// * Calculate spread between ask and bid
// * Have at least as much data to cover the amount of data that you have
// * Find the most recent candle MAX/MINIMA
// * find delta of highest high to lowest low

// If delta > or < than the interbal (RenkoLevel)
// The upper and lower bound is where you recalc the renko if the price moves across

type RealtimeContext struct {
	RenkoLevel int
	UpperBound int
	LowerBound int
	Direction  int
}

// var intervals = []int{10}

var intervals = []int{1000, 900, 800, 700, 600, 500, 400, 300, 200, 100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 5, 3, 2, 1}

func loadMockHistory() []data.FrontEndCandle {
	out := data.QueryCandleHistory()
	return out
}

/*
 */
func watchStream(prices chan data.Candle) {

	//Channel for candles coming from the api

	//Timer to query the api every six seconds
	ticker := time.NewTicker(3 * time.Second)

	//Exit signals to terminate the stream
	quit := make(chan struct{})
	e := make(chan int)

	//Instantiate renko array with default values of zero
	levels := buildRenkoLevels()

	// rawCandleHistory := data.QueryCandleHistory()
	// frontEndCandles := convertToFrontEndCandle(rawCandleHistory)

	frontEndCandles := data.QueryCandleHistory()

	renkoRanges := findLargestRenkoForHistory(frontEndCandles)
	//1.1250 ~ 0.96
	fmt.Println(renkoRanges)
	// bytes, err := json.Marshal(&frontEndCandles)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(bytes))

	// for _, candle := range frontEndCandles {
	// 	fmt.Printf("%v ,", candle)
	// }
	// data.QueryAskPrice()

	//Spawn a new thread that infinitely loops
	go func() {
		for {
			//Watch the channels for signals
			select {
			//When the timer emits a signal every six seconds . . .
			case <-ticker.C:
				// data.QueryAskPrice()
				//Get the latest 5 second candle from Oanda
				// rawCandleData := data.QueryCandleHistory()
				//Convert to a normalized form where the price value is in points (type int versus type float64)
				// candle := data.NewCandle(rawCandleData.Time, rawCandleData.Mid.O, rawCandleData.Mid.H, rawCandleData.Mid.L, rawCandleData.Mid.C)
				//TODO frontEndCandle := convertToFrontEndCandle(candles)

				//Send the new candle from Oanda to the prices channel
				// prices <- candle
			//Exit the program when the quit channel gets a signal
			case <-quit:
				e <- 1
				ticker.Stop()
				return
			}
		}
	}()
	//End new thread

	//Infinitely loop on this thread
	for {
		//Watch the channels named in the case statements
		select {
		//When the prices channel emits new data, give it a name of priceCandle as a local variable.
		//I.E: PriceCandle := (data coming out of <- ) prices
		case priceCandle := <-prices:
			//When a new priceCandle is received
			fmt.Printf("\n\n 5 Second Candle Closed @ \n %d", priceCandle.Close)
			//Use it to recalculate the renkos
			newContexts := compareRenkoLevels(priceCandle, levels)
			levels = newContexts
			fmt.Printf("\n %d, %v", newContexts[22].RenkoLevel, newContexts[22])
		//When exit signal received, return and exit thread
		case <-e:
			return
		}
	}

}

// Utiltiy function
func convertToFrontEndCandle(candles []models.CandleStick) []data.FrontEndCandle {
	newData := []data.FrontEndCandle{}
	for _, c := range candles {
		feC := data.NewFrontEndCandle(c.Time, c.Mid.O, c.Mid.H, c.Mid.L, c.Mid.C)
		newData = append(newData, feC)
	}
	return newData
}

// Build initial renkos with default values
func buildRenkoLevels() []RealtimeContext {
	contexts := []RealtimeContext{}
	for _, value := range intervals {
		contexts = append(contexts, RealtimeContext{RenkoLevel: value, UpperBound: 0, LowerBound: 0, Direction: 0})
	}
	return contexts
}

// Calculate the renkos on price action
// @param candle data.Candle - the new price action data
// @param levels []RealtimeContext - an array of all the renko information
func compareRenkoLevels(candle data.Candle, levels []RealtimeContext) []RealtimeContext {
	//Empty array
	newLevels := []RealtimeContext{}

	//Loop through the renko information that was passed in
	for _, context := range levels {

		//For convenience assign the current renko context to a variable
		renko := context.RenkoLevel

		//If this is the first pass, the values will be default of zero
		if context.UpperBound == 0 {

			//Set the upper and lower bound using the closing price
			context.UpperBound = renko + candle.Close
			context.LowerBound = candle.Close - renko

			//Add the newly calculated RealtimeContext into the array
			newLevels = append(newLevels, context)

			//If this is not the first pass, the values will not be zero
		} else {

			//If the closing price is above the upper bound, renko is moving up
			if candle.Close > context.UpperBound {
				context.Direction = 1
			}

			//If the closing price is below the lower bound, the renko is moving down
			if candle.Close < context.LowerBound {
				candle.Direction = -1
			}

			//Append the new renko context (RealtimeContext) into the array
			newLevels = append(newLevels, context)
		}
	}

	//Return the array
	return newLevels
}

func findLargestRenkoForHistory(h []data.FrontEndCandle) []float64 {
	output := []float64{}

	highest := data.FrontEndCandle{}
	lowest := data.FrontEndCandle{}

	fmt.Println(len(h))
	for i, c := range h {
		if i == 1 {
			highest = c
			lowest = c 
			continue
		}

		if c.High > highest.High {
			highest = c
		}

		if c.Low < lowest.Low { 
			lowest = c
		}
	}

	output = append(output, highest.High, lowest.Low)
	return output
}

var candlesMockHistory = []data.FrontEndCandle{}
