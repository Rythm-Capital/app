package old

// import (
// 	"testing"

// 	. "desktop.rythm/core/data"
// 	td "desktop.rythm/core/test"
// 	orderedmap "github.com/wk8/go-ordered-map/v2"
// )

// type ThresholdCandlestickTest struct {
// 	rythmCtx      RythmInitialContext
// 	glyph *orderedmap.OrderedMap[int,int]
// 	targetCandles map[int]Candle
// }

// func buildThreshCases(source map[int]map[string]interface{}) []ThresholdCandlestickTest {
	
// 	var slice []ThresholdCandlestickTest
	
// 	for renko, casedata := range source {

// 		testCase := ThresholdCandlestickTest{
// 			rythmCtx: initializeRythmContext(renko, casedata["candles"].(CandleHistory)),
// 			glyph: casedata["glyph"].(*orderedmap.OrderedMap[int,int]),
// 			targetCandles: casedata["expectedMatchingCandles"].(map[int]Candle),
// 		}

// 		slice = append(slice, testCase)
// 	}
	
// 	return slice
// }

// func Test_FindCandlesWithinThresholds(t *testing.T) {

// 	allCases := make(map[string][]ThresholdCandlestickTest)

// 	// chrono_down := buildThreshCases(td.ThresholdCandles_Chronological_Downtrend)
// 	trade_pattern_1 := buildThreshCases(td.ThresholdCandles_Trade_Pattern_1)

// 	// allCases["chrono_down"] = chrono_down
// 	allCases["trade_pattern_1"] = trade_pattern_1

	
// 	for testgroup, cases := range allCases {

// 		for _, tc := range cases {

// 			tc.rythmCtx.calculateRatioPrices(tc.glyph)
// 			tc.rythmCtx.findCandlesWithinThresholds(tc.glyph, 0)

// 			for key, targetCandle := range tc.targetCandles {
				
// 				matchedCandle, _ := tc.rythmCtx.candlesMatchingThresholds.Get(key)
				
// 				if matchedCandle != targetCandle {
// 					t.Errorf("\n\n %s \n Expected %v \n to match threshold at %d , got \n %v instead\n\n", testgroup, targetCandle, key, matchedCandle)
// 				}
// 			}

// 		}
// 	}
// }