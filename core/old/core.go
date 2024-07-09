package old

// import (

// 	// "sort"

// 	. "desktop.rythm/core/data"
// )

/* -------------------------------------------------------------------------- */
/*                                    TODO                                    */
/* -------------------------------------------------------------------------- */

//
// 1. The calculations need to take time data into account. See test file chronological examples
//		a. If you have two sets of candles with exact same prices, they could be either uptrending or
//		   downtrending based on their chronological order.
// //* Completed pending adverserial testing
//
//
// 2. Test coverage for all renko levels:
//	  1000, 990, 980, 970, 960, 950, 940, 930, 920, 910,
//	  900, 890, 880, 870, 860, 850, 840, 830, 820, 810,
//	  800, 790, 780, 770, 760, 750, 740, 730, 720, 710,
//	  700, 690, 680, 670, 660, 650, 640, 630, 620, 610,
//	  600, 590, 580, 570, 560, 550, 540, 530, 520, 510,
//	  500, 490, 480, 470, 460, 450, 440, 430, 420, 410,
//	  400, 390, 380, 370, 360, 350, 340, 330, 320, 310,
//	  300, 290, 280, 270, 260, 250, 240, 230, 220, 210,
//	  200, 190, 180, 170, 160, 150, 140, 130, 120, 110,
//	  100, 90, 80, 70, 60, 50, 40, 30, 20, 10
//
//
// 3. Need to generalize all forms to use any OHLC candle data, or the engine will miss trades.
//		a. All references to Candle.Close need to be expanded to use any candle data
// //* In progress
//
// 4. Chaotic test data beyond staple 10,20,30... need 17,81,31,12 etc
//
//
// 5. Case where multiple candles match the interval Cypher
//
// 6. Remove hard-coded dependency on an 'interval cypher'
//
//
//
//
//

//
//
//	--- Two functions
//	--- One , compare HighestCandle to LowestCandle,
//			if the price delta between the two is < a given renko,
//			that renko is UP
//
//  --- Two, compare LowestCandle to CurrentMarketPrice
//			 if the price delta between the two is is > a given renko,
//			 then that renko is UP
//
//  --- Chronological order the results of the two ^
//
//  --- Figure out a x-axis window and how it is bounded  (HH or LL become the new 'start')
//
//  ---

/* -------------------------------------------------------------------------- */
/*                                    Types                                   */
/* -------------------------------------------------------------------------- */

// type WavePhase struct {
// 	OpeningCandle Candle
// 	ClosingCandle Candle
// 	PhaseHeight   int
// 	Direction     int
// }

// type RythmInitialContext struct {
// 	history CandleHistory

// 	renko int

// 	highestCandle Candle
// 	lowestCandle  Candle
// 	direction     int

// 	ratioPriceValues          *orderedmap.OrderedMap[int, int]
// 	rythmPhases *orderedmap.OrderedMap[int, Candle]
// }

// TODO: Refactor
// type HotCandle struct {
// 	C             Candle
// 	PriceInterval int
// }

/* -------------------------------------------------------------------------- */
/*                            Init from chart data                            */
/* -------------------------------------------------------------------------- */

// func initializeRythmContext(renko int, h CandleHistory) RythmInitialContext {

// 	var highestCandle Candle
// 	var lowestCandle Candle
// 	var direction int

// 	for i, c := range h {

// 		if i == 0 {
// 			//First candle

// 			highestCandle = c
// 			lowestCandle = c

// 		} else {

// 			if c.High > highestCandle.High {
// 				highestCandle = c
// 			}
// 			if c.Low < lowestCandle.Low {
// 				lowestCandle = c
// 			}

// 		}

// 	}

// 	direction = CalculateDirection(lowestCandle, highestCandle, renko)

// 	fmt.Printf("\n\n INITIALIZING RYTHM CONTEXT:"+
// 		"\n Renko: %d"+
// 		"\n Highest Candle: %v"+
// 		"\n Lowest Candle: %v"+
// 		"\n Direction: %d \n", renko, highestCandle, lowestCandle, direction)

// 	return RythmInitialContext{history: h, renko: renko, highestCandle: highestCandle, lowestCandle: lowestCandle, direction: direction, rythmPhases: orderedmap.New[int, Candle]()}

// }

/* -------------------------------------------------------------------------- */
/*                                Rythm Context                               */
/* -------------------------------------------------------------------------- */

// func (rctx *RythmInitialContext) calculateRatioPrices(glyph *orderedmap.OrderedMap[int, int]) {

// 	var origin int
// 	priceIntervals := orderedmap.New[int, int]()

// 	//If the context is uptrending, use the highest high
// 	if rctx.direction == 1 {
// 		origin = rctx.highestCandle.High
// 	} else {
// 		//Otherwise, use the lowest low.
// 		origin = rctx.lowestCandle.Low
// 	}

// 	for glyphKey := glyph.Oldest(); glyphKey != nil; glyphKey = glyphKey.Next() {
// 		ratioAsDecimal := float64(glyphKey.Key) * 0.01
// 		priceIntervals.Set(glyphKey.Key, rctx.findRatioFrom(origin, ratioAsDecimal))
// 	}

// 	fmt.Printf("\n\n CALCULATING PRICE RATIOS BASED ON RECEIVED GLYPH"+
// 		"\n Glyph: %v"+
// 		"\n Calculated values: %v", glyph, priceIntervals)

// 	rctx.ratioPriceValues = priceIntervals
// }

// func (rctx *RythmInitialContext) findCandlesWithinThresholds(glyph *orderedmap.OrderedMap[int, int], searchIndex int) {

// 	//TODO: Needs to partition candle history into phases and delete them as a glyphKey is found
// 	// if original glyph - computational glyph = 0
// 	// searchIndex = 0

// 	computationalGlyph := *glyph

// 	glyphPhase := computationalGlyph.Oldest()
// 	glyphKey := glyphPhase.Key
// 	priceTarget := rctx.ratioPriceValues[glyphKey]
// 	intervalDirection := glyphPhase.Value

// 	fmt.Printf("\n\n Computational glyph: %v", computationalGlyph)
// 	fmt.Printf("\nSearch index: %d", searchIndex)
// 	fmt.Printf("\n Scanning candles for target price: %d for %d interval in direction: %d\n\n", priceTarget, glyphKey, intervalDirection)

// 	for i := searchIndex; i < len(rctx.history); i++ {

// 		candle := rctx.history[i]

// 		if intervalDirection == 1 {  // CONTEXT IS ASCENDING

// 			if candle.High == priceTarget { // CANDLE IS >= TARGET PRICE

// 				if rctx.candlesMatchingThresholds.Len() == 0 { // FIRST ITERATION, BASE CASE

// 					rctx.candlesMatchingThresholds.Set(glyphKey, candle)
// 					fmt.Printf("\n BASE CASE: Candle High above target price of %d found at: %v \n", priceTarget, candle)
// 					computationalGlyph.Delete(glyphPhase.Key)
// 					searchIndex = i
// 					break

// 				}

// 				if rctx.candlesMatchingThresholds.Newest().Value.Time.Before(candle.Time) { // CANDLE IS NEWER THAN PREVIOUS MATCH

// 					rctx.candlesMatchingThresholds.Set(glyphKey, candle)
// 					fmt.Printf("\n Candle High above target price of %d found at: %v \n", priceTarget, candle)
// 					computationalGlyph.Delete(glyphPhase.Key)
// 					searchIndex = i
// 					break

// 				}

// 			}

// 		} else if intervalDirection == -1 { // CONTEXT IS DESCENDING

// 			if candle.Low == priceTarget { // CANDLE IS <= TARGET PRICE

// 				if rctx.candlesMatchingThresholds.Len() == 0 { // FIRST ITERATION, BASE CASE
// 					rctx.candlesMatchingThresholds.Set(glyphKey, candle)
// 					fmt.Printf("\n BASE CASE: Candle Low below target price of %d found at: %v \n", priceTarget, candle)
// 					computationalGlyph.Delete(glyphPhase.Key)
// 					searchIndex = i
// 					break
// 				}

// 				if rctx.candlesMatchingThresholds.Newest().Value.Time.Before(candle.Time) { //CANDLE IS NEWER THAN PREVIOUS MATCH

// 					rctx.candlesMatchingThresholds.Set(glyphKey, candle)
// 					fmt.Printf("Previous match: %v",rctx.candlesMatchingThresholds.Oldest().Value)
// 					fmt.Printf("\n Candle Low below target price of %d found at: %v \n", priceTarget, candle)
// 					computationalGlyph.Delete(glyphPhase.Key)
// 					searchIndex = i
// 					break

// 				}

// 			}

// 		}

// 	}
// 	if computationalGlyph.Len() > 0 {
// 		rctx.findCandlesWithinThresholds(&computationalGlyph, searchIndex)
// 	} else {
// 		return
// 	}
// }

// func (rctx *RythmInitialContext) divideContextIntoPhases() []WavePhase {

// 	phases := []WavePhase{}

// 	return phases
// }

// TODO: HotCandle shit is hacky, find better way to sort
// func (rctx *RythmInitialContext) articulatePossibleTradePatterns() *orderedmap.OrderedMap[int, int] {

// 	hotCandles := []HotCandle{}

// 	tradeGlyph := orderedmap.New[int, int]()

// 	normalized_renko := rctx.renko / 10

// 	tradeGlyph.Set(normalized_renko, rctx.direction)

// 	for result := rctx.candlesMatchingThresholds.Oldest(); result != nil; result = result.Next() {
// 	// for priceInterval, candle := range rctx.candlesMatchingThresholds {
// 		hotCandles = append(hotCandles, HotCandle{C: result.Value, PriceInterval: result.Key})
// 	// }
// 	}

// 	sort.SliceStable(hotCandles, func(i, j int) bool {
// 		return hotCandles[i].C.Time.Before(hotCandles[j].C.Time)
// 	})

// 	for i, hotcandle := range hotCandles {

// 		if i == len(hotCandles)-1 {

// 			if hotcandle.PriceInterval < hotCandles[i].PriceInterval {
// 				tradeGlyph.Set(hotcandle.PriceInterval, 1)
// 			} else {
// 				tradeGlyph.Set(hotcandle.PriceInterval, -1)
// 			}

// 		} else {

// 			if hotcandle.PriceInterval < hotCandles[i+1].PriceInterval {
// 				tradeGlyph.Set(hotcandle.PriceInterval, 1)
// 			} else {
// 				tradeGlyph.Set(hotcandle.PriceInterval, -1)
// 			}

// 		}
// 	}

// 	// fmt.Printf("\n\nARTICULATING MATCHING TRADE PATTERNS" +
// 	// "\nArticulated Glyph: %v" +
// 	// "\n\n", tradeGlyph)

// 	return tradeGlyph
// }

/* -------------------------------------------------------------------------- */
/*                                    UTILS                                   */
/* -------------------------------------------------------------------------- */

// func (rctx RythmInitialContext) findRatioFrom(price int, ratio float64) int {
// 	sign := -rctx.direction //Sign is negative because it needs to be inverse of the direction of the context
// 	adjustor := sign * int(float64(rctx.renko)*ratio)
// 	calculatedValue := price + adjustor

// 	return calculatedValue
// }

// func CalculateDirection(lowestCandle Candle, highestCandle Candle, renko int) int {

// 	lowestLowInPoints := lowestCandle.Close
// 	highestHighInPoints := highestCandle.Close
// 	firstCandle := CompareTimestamps(lowestCandle, highestCandle)

// 	r := int(renko)

// 	if firstCandle.Time == highestCandle.Time {
// 		//The first candle is the highest, chart is going down

// 		if lowestLowInPoints < (highestHighInPoints - r) {
// 			//The lowest low is OUTSIDE renko
// 			return -1
// 		}
// 		//? If the price delta is WITHIN the renko, & the higher candle comes first...
// 		// ? Is the context descending?

// 	}

// 	//The first candle is the lowest, chart is going up
// 	//? If the price delta is OUTSIDE the renko, & the lowest candle comes first...
// 	//? Is the context ascending?

// 	return 1

// }

// func CompareTimestamps(a Candle, b Candle) (firstCandle Candle) {
// 	if a.Time.Before(b.Time) {
// 		firstCandle = a
// 		return
// 	} else {
// 		firstCandle = b
// 		return
// 	}
// }
