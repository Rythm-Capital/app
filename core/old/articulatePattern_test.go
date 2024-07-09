package old

// import (
// 	"testing"

// 	. "desktop.rythm/core/data"
// 	td "desktop.rythm/core/test"
// 	orderedmap "github.com/wk8/go-ordered-map/v2"
// )

// type InflectionTest struct {
// 	rythmCtx       RythmInitialContext
// 	intervalCypher []int
// 	targetGlyph    *orderedmap.OrderedMap[int, int]
// }

// func buildArtCases(source map[int]map[string]interface{}) []ArticulationTest {

// 	var slice []ArticulationTest

// 	for renko, casedata := range source {
// 		// fmt.Println(casedata["intervalCypher"])

// 		testCase := ArticulationTest{
// 			rythmCtx:       initializeRythmContext(renko, casedata["candles"].(CandleHistory)),
// 			intervalCypher: casedata["intervalCypher"].([]int),
// 			targetGlyph:    casedata["targetGlyph"].(*orderedmap.OrderedMap[int, int]),
// 		}

// 		slice = append(slice, testCase)

// 	}

// 	return slice
// }

// func Test_ArticulateTradePattern(t *testing.T) {

// 	allCases := make(map[string][]ArticulationTest)

// 	// chrono_down := buildArtCases(td.Articulation_Chronological_Downtrend)
// 	trade_pattern_1 := buildArtCases(td.Articulation_Trade_Pattern_1)

// 	// allCases["chrono_down"] = chrono_down
// 	allCases["trade_pattern_1"] = trade_pattern_1

// 	for testgroup, cases := range allCases {

// 		for _, tc := range cases {

// 			tc.rythmCtx.calculateRatioPrices(tc.intervalCypher)
// 			tc.rythmCtx.findCandlesWithinThresholds()
// 			resultingGlpyh := tc.rythmCtx.articulatePossibleTradePatterns()

// 			for result := resultingGlpyh.Oldest(); result != nil; result = result.Next() {
// 				for target := tc.targetGlyph.Oldest(); target != nil; target = target.Next() {

// 					if result.Key != target.Key {
// 						fmt.Printf("\n Articulation test for renko %s : \n Expected candle matching price interval %d. Got  %d instead \n\n", testgroup, target.Key, result.Key)
// 					}

// 					if result.Value != target.Value {
// 						fmt.Printf("\n Articulation test for renko %s : \n Expected candle direction matching price interval as %d. Got  %d instead \n\n", testgroup, target.Value, result.Value)
// 					}

// 				}
// 			}
// 		}
// 	}
// }

// func Test_FullLogging(t *testing.T) {

// 	allCases := make(map[string][]ArticulationTest)
// 	trade_pattern_1 := buildArtCases(td.Articulation_Trade_Pattern_1)

// 	allCases["trade_pattern_1"] = trade_pattern_1

// 	for _, cases := range allCases {

// 		for _, tc := range cases {
// 			tc.rythmCtx.calculateRatioPrices(tc.intervalCypher)
// 			tc.rythmCtx.findCandlesWithinThresholds()
// 			tc.rythmCtx.articulatePossibleTradePatterns()
// 		}
// 	}
// }
