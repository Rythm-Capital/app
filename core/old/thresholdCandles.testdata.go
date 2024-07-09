package old

// import (
// 	. "desktop.rythm/core/data"
// )


// var ThresholdCandles_Chronological_Downtrend = map[int]map[string]interface{}{
// 	1000: {
// 		"candles": chrono_down_candles,
// 		"glyph": TestGlyph,
// 		"expectedMatchingCandles": map[int]Candle{
// 			90: chrono_down_candles[4],
// 			// 70: chrono_down_candles[3],
// 			// 50: chrono_down_candles[2],
// 			// 10: chrono_down_candles[1],
// 		},
// 	},
// }

// //TODO: Chaotic test data 
// //TODO: Conversion utility between 1000,900,200 and normalized %s 
// var ThresholdCandles_Trade_Pattern_1 = map[int]map[string]interface{}{
// 	1000: {
// 		"candles": pattern1_candles,
// 		"glyph": GlyphOne_Normal,
// 		"expectedMatchingCandles": map[int]Candle{
// 			//[normalized percentage] : matching candle
// 			90: pattern1_candles[8],
// 			70: pattern1_candles[14],
// 			88: pattern1_candles[18],
// 			78: pattern1_candles[21],
// 			87: pattern1_candles[24],
// 		},
// 	},
// }