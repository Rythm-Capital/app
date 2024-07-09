package rythm

// import "testing"

import (
	// "testing"

	. "desktop.rythm/core/data"
	// td "desktop.rythm/core/test"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

//! WHAT IF: glyph is mismatched?
//! WHAT IF ending glyph key is not present ?
//! Chaotic Data / High noise data
//! Data with unnatural spikes
//! BAD data

type InflectionTest struct {
	h                     CandleHistory
	searchGlyph           *orderedmap.OrderedMap[int, int]
	targetInflectionGraph [][]int
}

func buildInflectionTestCases(testCases []map[string]interface{}) []InflectionTest {

	var slice []InflectionTest

	for _, casedata := range testCases {
        
		testCase := InflectionTest{
			h:                     casedata["h"].(CandleHistory),
			searchGlyph:           casedata["searchGlyph"].(*orderedmap.OrderedMap[int, int]),
			targetInflectionGraph: casedata["targetInflectionGraph"].([][]int),
		}

		slice = append(slice, testCase)

	}

	return slice
}

// func Test_Inflection(t *testing.T) {

// 	allCases := make(map[string][]InflectionTest)
	
// 	// simpleUptrendTests := buildInflectionTestCases(td.Simple_Uptrend_Tests)
// 	// simpleDowntrendTests := buildInflectionTestCases(td.Simple_Downtrend_Tests)

// 	// chronoUptrendTests := buildInflectionTestCases(td.Chronological_Uptrend_Tests)
// 	// chronoDowntrendTests := buildInflectionTestCases(td.Chronological_Downtrend_Tests)

// 	simpleWaveTests := buildInflectionTestCases(td.Simple_Wave_Tests)
	
// 	// allCases["Simple Uptrend Tests"] = simpleUptrendTests
// 	// allCases["Simple Downtrend Tests"] = simpleDowntrendTests

// 	// allCases["Chronological Uptrend Tests"] = chronoUptrendTests
// 	// allCases["Chronological Downtrend Tests"] = chronoDowntrendTests

// 	allCases["Simple Wave Tests"] = simpleWaveTests
	
// 	for testgroup, cases := range allCases {
		
// 		for _, tc := range cases {
			
// 			rctx := RythmContext{}
// 			rctx.InitializeRythmContext(tc.h, 0, tc.searchGlyph)
// 			resultingInflectionGraph := rctx.inflectionPoints

// 			if len(resultingInflectionGraph) != len(tc.targetInflectionGraph) {
// 				t.Fatalf("\n\n %s \n Inflection graph missed or exceeded target. \n Expected: %v \n Got: %v", testgroup, tc.targetInflectionGraph, resultingInflectionGraph )
// 			}

// 			for i, resultInflectionPointMetaData := range resultingInflectionGraph {

// 				for j, metaDataValues := range resultInflectionPointMetaData {

// 					if metaDataValues != tc.targetInflectionGraph[i][j] {
// 						t.Fatalf("\n\n %s: \n Expected %v, got %v", testgroup, tc.targetInflectionGraph, rctx.inflectionPoints)
// 					}

// 				}

// 			}
// 		}
// 	}
// }

// func Test_MisMatchedGlyph(t *testing.T) {

// 	rctx := RythmContext{}

// 	rctx.initializeRythmContext(td.Simple_quad_wave, 0, td.Simple_tri_Glyph)

// }
