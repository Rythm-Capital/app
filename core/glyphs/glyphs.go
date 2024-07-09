package glyphs

import (
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

// TODO: Explore methods of cross canceling redundant contexts,
// TODO: such as initial renko(maybe 90. .)
var TRADE_GLYPH_1 = [][]int{{1000, 900, 200, 180, 100, 90}, {1, -1, 1, -1, 1, -1}}
var normalized_TRADE_GLYPH_1 = [][]int{{100, 90, 70, 88, 78, 87}, {1, -1, 1, -1, 1, -1}}


type TradeLegend map[string]*orderedmap.OrderedMap[int, int]

var tradeGlyph_one = orderedmap.New[int, int]()
var tradeGlyph_one_normalized = orderedmap.New[int, int]()

func CreateTradeLegend() (TradeLegend) {
	tradeGlyph_one.Set(1000, -1)
	tradeGlyph_one.Set(900, -1)
	tradeGlyph_one.Set(200, 1)
	tradeGlyph_one.Set(180, -1)
	tradeGlyph_one.Set(100, 1)
	tradeGlyph_one.Set(90, -1)

	tradeGlyph_one_normalized.Set(100, 1)
	tradeGlyph_one_normalized.Set(90, -1)
	tradeGlyph_one_normalized.Set(70, 1)
	tradeGlyph_one_normalized.Set(88, -1)
	tradeGlyph_one_normalized.Set(78, 1)
	tradeGlyph_one_normalized.Set(87, -1)

	legend := make(map[string]*orderedmap.OrderedMap[int, int])
	
	legend["trade_gylph_one"] = tradeGlyph_one
	legend["trade_gylph_one_nm"] = tradeGlyph_one_normalized
	
	return legend
}
