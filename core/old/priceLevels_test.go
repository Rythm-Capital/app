package old

// import (
// 	"testing"

// 	. "desktop.rythm/core/data"
// 	td "desktop.rythm/core/test"
// 	orderedmap "github.com/wk8/go-ordered-map/v2"
// )

// type CalculateRatioPricesTestCase struct {
// 	rythmCtx     RythmInitialContext
// 	glyph        *orderedmap.OrderedMap[int, int]
// 	targetPrices *orderedmap.OrderedMap[int, []int]
// }

// func buildCalcCases(source map[int]map[string]interface{}) []CalculateRatioPricesTestCase {

// 	var slice []CalculateRatioPricesTestCase

// 	for renko, casedata := range source {

// 		testCase := CalculateRatioPricesTestCase{
// 			rythmCtx:     initializeRythmContext(renko, casedata["candles"].(CandleHistory)),
// 			glyph:        casedata["glyph"].(*orderedmap.OrderedMap[int, int]),
// 			targetPrices: casedata["expectedPriceLevels"].(*orderedmap.OrderedMap[int, []int]),
// 		}

// 		slice = append(slice, testCase)
// 	}

// 	return slice
// }

// func Test_CalculateRatioPrices(t *testing.T) {

// 	allCases := make(map[string][]CalculateRatioPricesTestCase)

// 	chrono_down := buildCalcCases(td.Levels_Chronological_Downtrend)
// 	trade_patttern_1 := buildCalcCases(td.Levels_Trade_Pattern_1)

// 	allCases["chrono_down"] = chrono_down
// 	allCases["trade_patttern_1"] = trade_patttern_1

// 	for testgroup, cases := range allCases {

// 		for _, tc := range cases {

// 			tc.rythmCtx.calculateRatioPrices(tc.glyph)
// 			for expectedPrices := tc.targetPrices.Oldest(); expectedPrices != nil; expectedPrices = expectedPrices.Next() {
				
// 				glyphKey := expectedPrices.Key
// 				priceAndDirection := expectedPrices.Value

// 				result, ok := tc.rythmCtx.ratioPriceValues.Get(glyphKey)
// 				if ok {
// 					if result != priceAndDirection[0] {
// 						t.Errorf("\n\n %s \n Renko Range of %d failed calculation for %d%% price level. \n Got %d, expected %d \n\n", testgroup, tc.rythmCtx.renko, glyphKey, result, priceAndDirection[0])
// 					}
// 				}
// 			}

// 		}
// 	}
// }
