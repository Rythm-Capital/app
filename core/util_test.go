package rythm

import (
	// "testing"
	// "time"

	// . "desktop.rythm/core/data"
	td "desktop.rythm/core/test"
)

var candles = td.Simple_Uptrend

// func Test_FindRatioFrom(t *testing.T) {
// 	rythmCtx := initializeRythmContext(1000, candles)
// 	ninetyPercentDrop := rythmCtx.findRatioFrom(ConvertToPoints(130.0), 0.9)

// 	if ninetyPercentDrop != ConvertToPoints(129.10) {
// 		t.Errorf("Find Ratio Calculation Faile Expected %d, Got %d", ConvertToPoints(129.10), ninetyPercentDrop)
// 	}
// }

// func Test_CalculateDirection(t *testing.T) {

// 	minuteInterval := time.Duration.Minutes(5)
// 	timeStep := time.Duration(minuteInterval)

// 	candleA := NewCandle(time.Now(), 129.10, 129.10, 129.10, 129.10)
// 	candleB := NewCandle(time.Now().Add(timeStep), 130.00, 130.00, 130.00, 130.00)

// 	direction := CalculateDirection(candleA, candleB, 1000)
// 	if direction != 1 {
// 		t.Errorf("Rythm Context Direction Calculation Faile Expected %d, Got %d", 1, direction)
// 	}
// }

// func Test_ConvertToPoints(t *testing.T) {
// 	result := ConvertToPoints(130.00)
// 	if result != 130000 {
// 		t.Errorf("Price Conversion to Points Faile Expected %d, Got %d", 130000, result)
// 	}
// }

// func Test_CreateNewCandle(t *testing.T) {

// 	candle := NewCandle(time.Now(), 130.00, 130.00, 130.00, 130.00)

// 	if candle.High != 130000 {
// 		t.Errorf("Candle did not convert float64 to interger point value.")
// 	}

// }

// func Test_CompareTimestamps(t *testing.T) {

// 	firstCandle := CompareTimestamps(candles[0], candles[1])

// 	if firstCandle.Time != candles[0].Time {
// 		t.Errorf("Timestamp comparison did not properly order candles. Expected %v, got %v", candles[0].Time, firstCandle.Time)
// 	}

// }
