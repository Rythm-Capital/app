package old

// import (
// 	. "desktop.rythm/core/data"
// 	orderedmap "github.com/wk8/go-ordered-map/v2"
// 	//orderedmap "github.com/wk8/go-ordered-map/v2"
// )

// func build_TargetsFor_Chrono_Down_1000() *orderedmap.OrderedMap[int, []int] {
// 	out := orderedmap.New[int, []int]()
// 	out.Set(90, []int{ConvertToPoints(129.10),-1})
// 	out.Set(80, []int{ConvertToPoints(129.20),-1})
// 	out.Set(70, []int{ConvertToPoints(129.30),-1})
// 	out.Set(60, []int{ConvertToPoints(129.40),-1})
// 	out.Set(50, []int{ConvertToPoints(129.50),-1})
// 	out.Set(40, []int{ConvertToPoints(129.60),-1})
// 	out.Set(30, []int{ConvertToPoints(129.70),-1})
// 	out.Set(20, []int{ConvertToPoints(129.80),-1})
// 	out.Set(10, []int{ConvertToPoints(129.90),-1})
// 	return out
// }

// func build_TargetsFor_Chrono_Down_100() *orderedmap.OrderedMap[int, []int] {
// 	out := orderedmap.New[int, []int]()
// 	out.Set(90, []int{ConvertToPoints(129.19),-1})
// 	out.Set(80, []int{ConvertToPoints(129.18),-1})
// 	out.Set(70, []int{ConvertToPoints(129.17),-1})
// 	out.Set(60, []int{ConvertToPoints(129.16),-1})
// 	out.Set(50, []int{ConvertToPoints(129.15),-1})
// 	out.Set(40, []int{ConvertToPoints(129.14),-1})
// 	out.Set(30, []int{ConvertToPoints(129.13),-1})
// 	out.Set(20, []int{ConvertToPoints(129.12),-1})
// 	out.Set(10, []int{ConvertToPoints(129.11),-1})
// 	return out
// }



// func build_TargetsFor_TradePattern_1() *orderedmap.OrderedMap[int, []int] {
// 	out := orderedmap.New[int, []int]()
// 	out.Set(90, []int{ConvertToPoints(129.10),-1})
// 	out.Set(70, []int{ConvertToPoints(129.30),1})
// 	out.Set(88, []int{ConvertToPoints(129.12),-1})
// 	out.Set(78, []int{ConvertToPoints(129.22),1})
// 	out.Set(87, []int{ConvertToPoints(129.13),-1})
// 	return out
// }

// var Levels_Chronological_Downtrend = map[int]map[string]interface{} {
// 	1000: {
// 		"candles": BuildCandleHistory(chronological_Downtrend_Candles),
// 		"glyph": TestGlyph,
// 		"expectedPriceLevels": build_TargetsFor_Chrono_Down_1000(),
// 	},
// 	100: {
// 		"candles": BuildCandleHistory(chronological_Downtrend_Candles),
// 		"glyph": TestGlyph,
// 		"expectedPriceLevels": build_TargetsFor_Chrono_Down_100(),			
// 		},
// }

// var Levels_Trade_Pattern_1 = map[int]map[string]interface{} {
// 	1000: {
// 		"candles": BuildCandleHistory(trade_Pattern1_candles_forcedHigh),
// 		"glyph": GlyphOne_Normal,
// 		"expectedPriceLevels": build_TargetsFor_TradePattern_1(),
// 	},
// }