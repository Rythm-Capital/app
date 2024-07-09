package testData

var Simple_Downtrend_Tests = []map[string]interface{}{
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_1000),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, -1} },
	},
	 {
		"h":   BuildCandleHistory(simple_downtrend_renko_1000),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_990),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {1, 990, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_950),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {5, 950, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_900),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {10, 900, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_890),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {11, 890, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_800),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {20, 800, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_700),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {30, 700, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_600),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {40, 600, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_500),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {50, 500, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_400),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {60, 400, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_300),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {70, 300, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_200),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {80, 200, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_100),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {90, 100, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_50),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {95, 50, -1} },
	},
	{
		"h":   BuildCandleHistory(simple_downtrend_renko_10),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, 1}, {99, 10, -1} },
	},
}