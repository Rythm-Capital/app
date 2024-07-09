package data

import (
	// "math"
	"time"
)

/* -------------------------------------------------------------------------- */
/*                                CANDLE TYPES                                */
/* -------------------------------------------------------------------------- */

type Candle struct {
	Time      time.Time `json:"date"`
	Open      int       `json:"open"`
	High      int       `json:"high"`
	Low       int       `json:"low"`
	Close     int       `json:"close"`
	Direction int       `json:"direction"`
	// volume int
}

type FrontEndCandle struct {
	Time      time.Time `json:"date"`
	Open      float64       `json:"open"`
	High      float64       `json:"high"`
	Low       float64       `json:"low"`
	Close     float64       `json:"close"`
	// volume int
}

func NewCandle(time time.Time, open float64, high float64, low float64, close float64) Candle {
	direction := 0

	if open > close {
		direction = -1
	} else {
		direction = 1
	}

	return Candle{time, ConvertToPoints(open), ConvertToPoints(high), ConvertToPoints(low), ConvertToPoints(close), direction}
}

type CandleHistory []Candle

/* -------------------------------------------------------------------------- */
/*                                 CONVERSIONS                                */
/* -------------------------------------------------------------------------- */

func ConvertToPoints(number float64) int {
	return int(number / 0.00001)
}

func NormalizeRenko(renko int) int {
	return 0
}

func NewFrontEndCandle(t time.Time, open float64, high float64, low float64, close float64) FrontEndCandle {

	return FrontEndCandle{t, open, high, low, close}
}
