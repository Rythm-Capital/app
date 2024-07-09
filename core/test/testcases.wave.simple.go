package testData

var Simple_Wave_Tests = []map[string]interface{}{
	{
		"h":                     Simple_Wave,
		"searchGlyph":           Generic_Glyph,
		"targetInflectionGraph": [][]int{{0, 1000, 1}, {10, 900, -1}, {40, 600, 1} },
	},
	{
		"h":                     Simple_Tri_Wave,
		"searchGlyph":           Simple_Tri_Glyph,
		"targetInflectionGraph": [][]int{{0, 1000, 1}, {10, 900, -1}, {40, 600, 1}, {36, 640, -1}},
	},
	{
		"h":                     Simple_Quad_Wave,
		"searchGlyph":           Simple_Quad_Glyph,
		"targetInflectionGraph": [][]int{{0, 1000, 1}, {10, 900, -1}, {40, 600, 1}, {36, 640, -1}, {60, 400, 1}},
	},
}

var Mismatched_Glyph_Tests = []map[string]interface{}{
	// {
	// 	"h": Simple_Quad_Wave,
	// 	"searchGlyph": Simple_Tri_Glyph,
	// "targetInflectionGraph": [][]int{}
	// }
}
