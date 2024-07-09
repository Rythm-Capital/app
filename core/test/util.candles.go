package testData

import (
	"time"

	. "desktop.rythm/core/data"
)

/* -------------------------------------------------------------------------- */
/*                            CANDLE HISTORY UTILS                            */
/* -------------------------------------------------------------------------- */

type RawCandle struct {
	open  float64
	high  float64
	low   float64
	close float64
	// volume float64
}

func BuildCandleHistory(items []RawCandle) CandleHistory {

	var h CandleHistory

	for i, rawData := range items {

		interval := 5 * i
		minuteInterval := time.Duration(interval) * time.Minute

		candle := NewCandle(
			time.Now().Add(minuteInterval),
			rawData.open,
			rawData.high,
			rawData.low,
			rawData.close,
		)

		h = append(h, candle)
	}

	return h

}


func BuildFrontEndCandles(items []RawCandle) []FrontEndCandle {
	s := []FrontEndCandle{}
	
	for i, rawData := range items {
		
		minuteInterval := 5 * i
		timeStep := time.Duration(minuteInterval) * time.Minute
		candle := NewFrontEndCandle(time.Now().Add(timeStep), rawData.open, rawData.high, rawData.low, rawData.close)

		s = append(s, candle)
	}

	return s
}

// /* -------------------------------------------------------------------------- */
// /*                              TEST EXPECTATIONS                             */
// /* -------------------------------------------------------------------------- */

// var Chronological_Uptrend = map[int]map[string]interface{}{
// 	10000: {
// 		"history":   BuildCandleHistory(chronological_uptrend_candles),
// 		"high":      ConvertToPoints(130.00),
// 		"low":       ConvertToPoints(129.10),
// 		"direction": 1,
// 	},
// 	1: {
// 		"history":   BuildCandleHistory(chronological_uptrend_candles),
// 		"high":      ConvertToPoints(130.00),
// 		"low":       ConvertToPoints(129.10),
// 		"direction": 1,
// 	},
// }

