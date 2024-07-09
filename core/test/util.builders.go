package testData

import orderedmap "github.com/wk8/go-ordered-map/v2"


/* -------------------------------------------------------------------------- */
/*                                  FRONT-END                                 */
/* -------------------------------------------------------------------------- */

var Simple_Wave_Frontend = BuildFrontEndCandles(simple_wave_2)

/* -------------------------------------------------------------------------- */
/*                              CANDLE HISTORIES                              */
/* -------------------------------------------------------------------------- */

var Simple_Uptrend = BuildCandleHistory(simple_uptrend_renko_1000)

var Simple_Wave = BuildCandleHistory(simple_wave)
var Simple_Tri_Wave = BuildCandleHistory(simple_tri_wave)
var Simple_Quad_Wave = BuildCandleHistory(simple_quad_wave)

var Pattern1_candles = BuildCandleHistory(trade_Pattern1_candles_forcedHigh)

/* -------------------------------------------------------------------------- */
/*                                   GLYPHS                                   */
/* -------------------------------------------------------------------------- */

var Generic_Glyph = build_Generic_Glyph()

var Simple_Descending_Glyph = build_SimpleDescending_Glyph()

var Simple_Tri_Glyph = build_Tri_Glyph()
var Simple_Quad_Glyph = build_Quad_Glyph()

var TradePattern1_Glyph = build_TradePattern1_Glyph()

/* -------------------------------------------------------------------------- */
/*                                    UTILS                                   */
/* -------------------------------------------------------------------------- */

func build_Generic_Glyph() *orderedmap.OrderedMap[int, int] {
	testGlyph := orderedmap.New[int, int]()
	testGlyph.Set(0, 0)
	return testGlyph
}

func build_SimpleDescending_Glyph() *orderedmap.OrderedMap[int, int] {
	testGlyph := orderedmap.New[int, int]()
	testGlyph.Set(900, -1)
	testGlyph.Set(800, 1)
	testGlyph.Set(700, 1)
	testGlyph.Set(600, 1)
	testGlyph.Set(500, 1)
	testGlyph.Set(400, 1)
	testGlyph.Set(300, 1)
	testGlyph.Set(200, 1)
	testGlyph.Set(100, 1)
	return testGlyph
}

func build_Tri_Glyph() *orderedmap.OrderedMap[int, int] {
	testGlyph := orderedmap.New[int, int]()
	testGlyph.Set(1000, 1)
	testGlyph.Set(900, -1)
	testGlyph.Set(600, 1)			
	testGlyph.Set(640, -1) 
	return testGlyph
}

func build_Quad_Glyph() *orderedmap.OrderedMap[int, int] {
	testGlyph := orderedmap.New[int, int]()
	testGlyph.Set(1000, 1)
	testGlyph.Set(900, -1)
	testGlyph.Set(600, 1)
	testGlyph.Set(640, -1)
	testGlyph.Set(400, 1)
	return testGlyph
}

func build_TradePattern1_Glyph() *orderedmap.OrderedMap[int, int] {
	testGlyph := orderedmap.New[int, int]()
	testGlyph.Set(100, 1)
	testGlyph.Set(90, -1)
	testGlyph.Set(70, 1)
	testGlyph.Set(88, -1)
	testGlyph.Set(78, 1)
	testGlyph.Set(87, -1)
	return testGlyph
}



