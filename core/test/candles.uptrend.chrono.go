package testData


var chronological_uptrend_candles = []RawCandle{
	{129.10, 129.10, 129.10, 129.10},
	{129.30, 129.30, 129.30, 129.30},
	{129.50, 129.50, 129.50, 129.50},
	{129.90, 129.90, 129.90, 129.90},
	{130.00, 130.00, 130.00, 130.00},
}
 
var Chronological_Uptrend_Tests = []map[string]interface{}{
	{
		"h":   BuildCandleHistory(chronological_uptrend_candles),
		"searchGlyph": Generic_Glyph,
		"targetInflectionGraph": [][]int{ {0, 1000, -1}, {10, 900, 1} },
	},
}