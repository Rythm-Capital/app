package rythm

import (
	"fmt"

	. "desktop.rythm/core/data"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

/* -------------------------------------------------------------------------- */
/*                                    Types                                   */
/* -------------------------------------------------------------------------- */

type WavePhase struct {
	OpeningCandle Candle
	ClosingCandle Candle
	Direction     int
	PhaseRenkoMap *orderedmap.OrderedMap[int, int]
}

type RythmContext struct {
	phaseHistory     []WavePhase
	inflectionPoints [][]int
}

func NewRythmContext() *RythmContext {
	return &RythmContext{}
}

/* -------------------------------------------------------------------------- */
/*                            Init from chart data                            */
/* -------------------------------------------------------------------------- */

/*
InitializeRythmContext

	@param History: A given historical chart of candle data.
	@param SearchIndex: Recursive parameter, manually pass 0 (zero) at call site.
	@param TradeGlyph: An ordered map of type [int, int] that holds target renko values paired with their corresponding directionality.

	InitalizeRythmContext is a recursive function, where the exit case is hit when the search index is equal to the last index in the history.
	Upon the exit case being hit, a final function call will be made to detect the renko value and directionality of the current market price as it fits within the trade glyph.

	On first pass, the algorithim will execute the following operations:

	1. Matches the highest and lowest candles in the history.
	2. Passes the highest and lowest candle to PhaseDirection(), which will order them and return a directionality.
	3. Create a new WavePhase struct, to hold metadata about this pass for later use.
	4. From the WavePhase, find the renko deltas. This will loop through all renko values and assign 1 or -1 if the WavePhase is ascending or descending in that range.
	5. Find the inflection point of the renko map. The renko map will follow a pattern where all values are either 1 or -1 and inflect at a certain value. This value is found and returned.
	6. Append the WavePhase and Inflection values to an array.
	7. Make a recursive call to self with a new search index, effectively sliding the x-axis window to the right.

	On all following subsequent calls to self:

	1(a). Matches any candles that fall within one cents (0.01) of a pattern target level, OR: the highest and lowest candles given that no significance is otherwise found within price movement.
*/
func (rctx *RythmContext) InitializeRythmContext(h CandleHistory, searchIndex int, tradeGlyph *orderedmap.OrderedMap[int, int]) {

	//Exit Case
	if searchIndex == len(h)-1 {
		fmt.Printf("\n\n Inflection graph: ")
		for _, inf := range rctx.inflectionPoints {
			fmt.Printf("\n %d : %d ", inf[1], inf[2])
		}
		fmt.Printf("\n\n")
		//TODO: Compare to current market price for final 'trigger' event
		return
	}

	//Variables
	var highestCandle Candle
	var indexHighestCandle int

	var lowestCandle Candle
	var indexLowestCandle int

	var direction int

	var priceBounceOffset int
	var priceBounceDirection int

	//The current index of the inflection graph as it is being recursively built out.
	currentInflectionIndex := len(rctx.inflectionPoints) - 1

	//If this is the first pass, set values to zero in order to short-circuit later conditional checks.
	if currentInflectionIndex < 0 {

		currentInflectionIndex = 0
		priceBounceOffset = 0
		priceBounceDirection = 0

	} else {

		//Otherwise, retrieve the last recorded item from the inflection graph, of form [][]int. IE: 900, 600, etc.
		lastMeaningfulInterval := rctx.inflectionPoints[currentInflectionIndex][1]

		//In the TradeGlyph, retrieve the paired values that correspond to the last recorded item from the inflection graph.
		glyph := tradeGlyph.GetPair(lastMeaningfulInterval)

		if glyph != nil {

			next := glyph.Next() //Look ahead for what the next meaningful price action will be.

			if next != nil {
				priceBounceOffset = next.Key
				priceBounceDirection = next.Value
			}
			// and assign it to memory in order to dynamically set price boundaries to detect movement.
		}
	}

	//Open our search with the candle that matches our x-axis search index, 0 on first pass.
	phaseOpeningCandle := h[searchIndex]

	fmt.Printf("\n\n *** *** Recursive pass begining with search index: %d"+
		"\n\n Opening with search at candle %v \n", searchIndex, phaseOpeningCandle)

	for i := searchIndex; i < len(h); i++ {

		if i == searchIndex {

			//Within the loop, set these memory positions to the first candle, as a basis for calculation.
			highestCandle = h[i]
			lowestCandle = h[i]

		} else { // The second and subsequent iterations of i

			// If the price bounce direction is not zero (i.e we are in a current recursive call to self), then
			// set artificial boundary conditions within a search block in order to check
			// if any movement matches the pattern given in the trade glyph.
			if priceBounceDirection == 1 {

				if rctx.priceBouncesWithinTheshold(h[i], priceBounceOffset, priceBounceDirection) {

					//If the candles bounces here while we are ascending, then this was the highest candle.
					highestCandle = h[i]
					indexHighestCandle = i
					break
				}

			} else if priceBounceDirection == -1 {

				if rctx.priceBouncesWithinTheshold(h[i], priceBounceOffset, priceBounceDirection) {

					//If the candle bounces here while we are descending, then this was the lowest candle.
					lowestCandle = h[i]
					indexLowestCandle = i
					break
				}

			}

			//If first pass or no meaningful price action found, use simply the tallest or lowest candle.
			if h[i].High >= highestCandle.High {
				highestCandle = h[i]
				indexHighestCandle = i
				//Save index of candle found
			}
			if h[i].Low <= lowestCandle.Low {
				lowestCandle = h[i]
				indexLowestCandle = i
				//Save index of candle found
			}

		}

	}

	fmt.Println("\n Found Highest Candle:", highestCandle)
	fmt.Println(" Found Lowest Candle:", lowestCandle)

	//Order candles
	firstCandle, secondCandle, direction := PhaseDirection(lowestCandle, highestCandle)

	//Create memory structure
	phase := WavePhase{firstCandle, secondCandle, direction, orderedmap.New[int, int]()}

	//! extra guarding ? -- may help truncate recursion if we find ourselves in an edge cases at the end of a history. TODO
	// if phase.Direction == 1 {
	// 	if (indexHighestCandle == searchIndex) {
	// 		return
	// 	}
	// } else if phase.Direction == -1 {
	// 	if indexLowestCandle == searchIndex {
	// 		return
	// 	}
	// }

	fmt.Printf("\n\n CREATED A WAVE PHASE WITH PROPERTIES:"+
		"\n Opening Candle: %v"+
		"\n Closing Candle: %v"+
		"\n Direction: %d \n\n", firstCandle, secondCandle, direction)

	//Calculate renko deltas for the newly instantiated phase
	phase.findRenkoDeltas(rctx.phaseHistory)

	//Append opening renko value if this is the first phase
	if len(rctx.phaseHistory) < 1 {

		outerIntervalDirection, ok := phase.PhaseRenkoMap.Get(1000)

		if ok {
			rctx.inflectionPoints = append(rctx.inflectionPoints, []int{0, 1000, outerIntervalDirection})
		}

	}

	//Find the inflection point
	inflection := phase.findInflectionPoint()

	//Record the index of the renko map where the inflection point was found, the inflection point value, and the direction of the phase.
	
	if inflection[0] != 0 { // Ensure the inflection index is not zero: i.e.: that it exists, or that an inflection point /was/ found.
	
		rctx.inflectionPoints = append(rctx.inflectionPoints, []int{inflection[0], inflection[1], phase.Direction})
	
	}

	//Append the phase into a memory location
	rctx.phaseHistory = append(rctx.phaseHistory, phase)

	fmt.Printf("\n WAVE INFLECTION GRAPH DATA:"+
		"\n Inflection Index: %d"+
		"\n Inflection Renko Value: %d"+
		"\n Inflection Direction: %d", inflection[0], inflection[1], phase.Direction)

	if phase.Direction == 1 {
		//If this was an ascending phase (low to high), then recursively call self with the search index moved to the index where the highest candle was found.
		rctx.InitializeRythmContext(h, indexHighestCandle, tradeGlyph)
	} else if phase.Direction == -1 {
		//If this was an descending phase (high to low), then recursively call self with the search index moved to the index where the lowest candle was found.
		rctx.InitializeRythmContext(h, indexLowestCandle, tradeGlyph)
	}

	//Function repeats with the search index moved to the right along the x-axis. Effectively draws a rectangle along the chart with the Y set by high/low price delta and X set by high/low time delta.
	//Upon repeating, high/low pairs that have already been matched are ignored (search window slides to the right).
} 
//:END MAIN 

// Calculates a price boundary and returns true if a candle has fallen within one cent (0.01) of the target.
func (rctx *RythmContext) priceBouncesWithinTheshold(candle Candle, priceOffsetDelta int, priceBounceDirection int) bool {

	//The first wave phase that was computed
	outmostReferencePhase := rctx.phaseHistory[0]
	outmostDirection := outmostReferencePhase.Direction

	//The highest and lowest candles in the given chart history
	referenceOpenCandle := outmostReferencePhase.OpeningCandle
	referenceCloseCandle := outmostReferencePhase.ClosingCandle

	var priceBounceLine int

	if outmostDirection == -1 { //If the first wave phase was descending

		//Then a price boundary is set above the opening (highest) candle by subtracting the offset delta.
		priceBounceLine = referenceOpenCandle.High - priceOffsetDelta

	} else if outmostDirection == 1 { //If the first wave phase was ascending

		//Then a price boundary is set above the closing (lowest) candle by adding the offset delta.
		priceBounceLine = referenceCloseCandle.Low + priceOffsetDelta

	}

	if priceBounceDirection == 1 { //If the wave phase we are currently computing is ascending

		fmt.Printf("\n + + + Setting price bounce line at %d \n", priceBounceLine)

		//Return true if the candle's HIGH falls within one cent (0.01) of the bounce line.
		if candle.High >= priceBounceLine-10 && candle.High <= priceBounceLine+10 { //TODO: Compare open, close?
			fmt.Printf("\n !! Price bounced at %v\n", candle)
			return true
		}

	} else if priceBounceDirection == -1 { //If the wave phase we are currently computing is descending

		fmt.Printf("\n - - - Setting price bounce line at %d \n", priceBounceLine)

		//Return true if the candle's LOW falls within one cent (0.01) of the bounce line.
		if candle.Low >= priceBounceLine-10 && candle.Low <= priceBounceLine+10 { //TODO: Compare open, close?
			fmt.Printf("\n !! Price bounced at %v\n", candle)
			return true
		}

	}

	fmt.Println("\n Candle did not bounce.")

	//Return false if the candle did not match the bounce line tolerance.
	return false
}

// Compares the timestamps of the candles passed and returns them in chronological order along with a directionality of 1 (ascending) or -1 (descending).
func PhaseDirection(lowestCandle Candle, highestCandle Candle) (firstCandle Candle, secondCandle Candle, direction int) {
	if lowestCandle.Time.Before(highestCandle.Time) {
		firstCandle = lowestCandle
		secondCandle = highestCandle
		direction = 1
		return
	} else {
		firstCandle = highestCandle
		secondCandle = lowestCandle
		direction = -1
		return
	}
}

/*
Takes the price delta between the highest and lowest candle and uses that to determine whether the corresponding renko interval is ascending or descending.
The first wave phase sets the tallest price delta, so all subsequent wave phases must normalize their calculations of renko to align with the initially defined phase.
*/
func (w *WavePhase) findRenkoDeltas(phaseHistory []WavePhase) {
	var delta int

	//If there is already a wave phase in memory,
	if len(phaseHistory) > 0 {

		outMostPhaseContext := phaseHistory[0]

		referenceOpenCandle := outMostPhaseContext.OpeningCandle
		referenceCloseCandle := outMostPhaseContext.ClosingCandle

		if outMostPhaseContext.Direction == -1 { //If our entire context, initial phase was descending,

			//Compare the initial Open Candle (the highest candle found in the history) to the current candle on the end of the phase to get the price delta
			//TODO: Compare the current wave phase to determine if to use w.Opening or w.Closing
			delta = referenceOpenCandle.High - w.ClosingCandle.High

		} else if outMostPhaseContext.Direction == 1 { //If the entire context, initial phase was ascending,

			//Compare the initial Close Candle (the highest candle found in the history) to the current candle on the end of the phase to get the price delta.
			//TODO: Compare the current wave phase to determine if to use w.Opening or w.Closing
			delta = referenceCloseCandle.High - w.ClosingCandle.Low

		}

	} else { //If this is the first phase being computed.

		if w.Direction == -1 { //And we are descending,

			// Compare the opening candle (the highest candle found in the history) to the closing (lowest) candle in the phase to get the price delta.
			delta = w.OpeningCandle.High - w.ClosingCandle.Low

		} else if w.Direction == 1 {

			// Compare the closing candle (the highest candle found in the history) to the opening (lowest) candle in the phase to get the price delta.
			delta = w.ClosingCandle.High - w.OpeningCandle.Low

		}
	}

	//Create empty renko map
	renkoMap := InitRenkoMap()

	//Loop through all values
	for renkoEntry := renkoMap.Oldest(); renkoEntry != nil; renkoEntry = renkoEntry.Next() {

		if w.Direction == -1 { //If context is descending

			if delta < renkoEntry.Key { //And delta is below the renko entry,

				fmt.Println(renkoEntry.Key, ": 1")
				renkoEntry.Value = 1
				//We are ascending in this range

			} else if delta >= renkoEntry.Key { //If delta is above the renko entry

				fmt.Println(renkoEntry.Key, ": -1")

				renkoEntry.Value = -1
				//We are descending in this range

			}

		} else if w.Direction == 1 { //If context is ascending

			if delta >= renkoEntry.Key { //And delta is above the renko entry

				fmt.Println(renkoEntry.Key, ": 1")
				renkoEntry.Value = 1
				//We are ascending in this range

			} else if delta < renkoEntry.Key { //If delta is below renko entry

				fmt.Println(renkoEntry.Key, ": -1")
				renkoEntry.Value = -1
				//We are descending in this range
			}

		}

	}

	//Save map into memory for the given phase.
	w.PhaseRenkoMap = renkoMap

}

/*
Iterates through the renko map and finds the points where the two values sum to zero, IE where the first value is -1 and the next 1.
Returns the index of any such point along with the renko interval. []int{index, renkoInterval}
*/
func (w *WavePhase) findInflectionPoint() []int {
	i := 0
	for r := w.PhaseRenkoMap.Oldest(); r != nil; r = r.Next() {
		if r.Next() != nil {
			if r.Value+r.Next().Value == 0 {
				i++
				return []int{i, r.Next().Key}
			}
		}
		i++
	}
	return []int{0, 0} // escape case
}



// Creates an ordered map with all ranges of renko values
func InitRenkoMap() *orderedmap.OrderedMap[int, int] {
	renkoMap := orderedmap.New[int, int]()

	a := orderedmap.Pair[int, int]{Key: 1000, Value: 0}
	b := orderedmap.Pair[int, int]{Key: 990, Value: 0}
	c := orderedmap.Pair[int, int]{Key: 980, Value: 0}
	d := orderedmap.Pair[int, int]{Key: 970, Value: 0}
	e := orderedmap.Pair[int, int]{Key: 960, Value: 0}
	f := orderedmap.Pair[int, int]{Key: 950, Value: 0}
	g := orderedmap.Pair[int, int]{Key: 940, Value: 0}
	h := orderedmap.Pair[int, int]{Key: 930, Value: 0}
	i := orderedmap.Pair[int, int]{Key: 920, Value: 0}
	j := orderedmap.Pair[int, int]{Key: 910, Value: 0}
	k := orderedmap.Pair[int, int]{Key: 900, Value: 0}
	l := orderedmap.Pair[int, int]{Key: 890, Value: 0}
	m := orderedmap.Pair[int, int]{Key: 880, Value: 0}
	n := orderedmap.Pair[int, int]{Key: 870, Value: 0}
	o := orderedmap.Pair[int, int]{Key: 860, Value: 0}
	p := orderedmap.Pair[int, int]{Key: 850, Value: 0}
	q := orderedmap.Pair[int, int]{Key: 840, Value: 0}
	r := orderedmap.Pair[int, int]{Key: 830, Value: 0}
	s := orderedmap.Pair[int, int]{Key: 820, Value: 0}
	t := orderedmap.Pair[int, int]{Key: 810, Value: 0}
	u := orderedmap.Pair[int, int]{Key: 800, Value: 0}
	v := orderedmap.Pair[int, int]{Key: 790, Value: 0}
	w := orderedmap.Pair[int, int]{Key: 780, Value: 0}
	x := orderedmap.Pair[int, int]{Key: 770, Value: 0}
	y := orderedmap.Pair[int, int]{Key: 760, Value: 0}
	z := orderedmap.Pair[int, int]{Key: 750, Value: 0}

	a1 := orderedmap.Pair[int, int]{Key: 740, Value: 0}
	b1 := orderedmap.Pair[int, int]{Key: 730, Value: 0}
	c1 := orderedmap.Pair[int, int]{Key: 720, Value: 0}
	d1 := orderedmap.Pair[int, int]{Key: 710, Value: 0}
	e1 := orderedmap.Pair[int, int]{Key: 700, Value: 0}
	f1 := orderedmap.Pair[int, int]{Key: 690, Value: 0}
	g1 := orderedmap.Pair[int, int]{Key: 680, Value: 0}
	h1 := orderedmap.Pair[int, int]{Key: 670, Value: 0}
	i1 := orderedmap.Pair[int, int]{Key: 660, Value: 0}
	j1 := orderedmap.Pair[int, int]{Key: 650, Value: 0}
	k1 := orderedmap.Pair[int, int]{Key: 640, Value: 0}
	l1 := orderedmap.Pair[int, int]{Key: 630, Value: 0}
	m1 := orderedmap.Pair[int, int]{Key: 620, Value: 0}
	n1 := orderedmap.Pair[int, int]{Key: 610, Value: 0}
	o1 := orderedmap.Pair[int, int]{Key: 600, Value: 0}
	p1 := orderedmap.Pair[int, int]{Key: 590, Value: 0}
	q1 := orderedmap.Pair[int, int]{Key: 580, Value: 0}
	r1 := orderedmap.Pair[int, int]{Key: 570, Value: 0}
	s1 := orderedmap.Pair[int, int]{Key: 560, Value: 0}
	t1 := orderedmap.Pair[int, int]{Key: 550, Value: 0}
	u1 := orderedmap.Pair[int, int]{Key: 540, Value: 0}
	v1 := orderedmap.Pair[int, int]{Key: 530, Value: 0}
	w1 := orderedmap.Pair[int, int]{Key: 520, Value: 0}
	x1 := orderedmap.Pair[int, int]{Key: 510, Value: 0}
	y1 := orderedmap.Pair[int, int]{Key: 500, Value: 0}
	z1 := orderedmap.Pair[int, int]{Key: 490, Value: 0}

	a2 := orderedmap.Pair[int, int]{Key: 480, Value: 0}
	b2 := orderedmap.Pair[int, int]{Key: 470, Value: 0}
	c2 := orderedmap.Pair[int, int]{Key: 460, Value: 0}
	d2 := orderedmap.Pair[int, int]{Key: 450, Value: 0}
	e2 := orderedmap.Pair[int, int]{Key: 440, Value: 0}
	f2 := orderedmap.Pair[int, int]{Key: 430, Value: 0}
	g2 := orderedmap.Pair[int, int]{Key: 420, Value: 0}
	h2 := orderedmap.Pair[int, int]{Key: 410, Value: 0}
	i2 := orderedmap.Pair[int, int]{Key: 400, Value: 0}
	j2 := orderedmap.Pair[int, int]{Key: 390, Value: 0}
	k2 := orderedmap.Pair[int, int]{Key: 380, Value: 0}
	l2 := orderedmap.Pair[int, int]{Key: 370, Value: 0}
	m2 := orderedmap.Pair[int, int]{Key: 360, Value: 0}
	n2 := orderedmap.Pair[int, int]{Key: 350, Value: 0}
	o2 := orderedmap.Pair[int, int]{Key: 340, Value: 0}
	p2 := orderedmap.Pair[int, int]{Key: 330, Value: 0}
	q2 := orderedmap.Pair[int, int]{Key: 320, Value: 0}
	r2 := orderedmap.Pair[int, int]{Key: 310, Value: 0}
	s2 := orderedmap.Pair[int, int]{Key: 300, Value: 0}
	t2 := orderedmap.Pair[int, int]{Key: 290, Value: 0}
	u2 := orderedmap.Pair[int, int]{Key: 280, Value: 0}
	v2 := orderedmap.Pair[int, int]{Key: 270, Value: 0}
	w2 := orderedmap.Pair[int, int]{Key: 260, Value: 0}
	x2 := orderedmap.Pair[int, int]{Key: 250, Value: 0}
	y2 := orderedmap.Pair[int, int]{Key: 240, Value: 0}
	z2 := orderedmap.Pair[int, int]{Key: 230, Value: 0}

	a3 := orderedmap.Pair[int, int]{Key: 220, Value: 0}
	b3 := orderedmap.Pair[int, int]{Key: 210, Value: 0}
	c3 := orderedmap.Pair[int, int]{Key: 200, Value: 0}
	d3 := orderedmap.Pair[int, int]{Key: 190, Value: 0}
	e3 := orderedmap.Pair[int, int]{Key: 180, Value: 0}
	f3 := orderedmap.Pair[int, int]{Key: 170, Value: 0}
	g3 := orderedmap.Pair[int, int]{Key: 160, Value: 0}
	h3 := orderedmap.Pair[int, int]{Key: 150, Value: 0}
	i3 := orderedmap.Pair[int, int]{Key: 140, Value: 0}
	j3 := orderedmap.Pair[int, int]{Key: 130, Value: 0}
	k3 := orderedmap.Pair[int, int]{Key: 120, Value: 0}
	l3 := orderedmap.Pair[int, int]{Key: 110, Value: 0}
	m3 := orderedmap.Pair[int, int]{Key: 100, Value: 0}
	n3 := orderedmap.Pair[int, int]{Key: 90, Value: 0}
	o3 := orderedmap.Pair[int, int]{Key: 80, Value: 0}
	p3 := orderedmap.Pair[int, int]{Key: 70, Value: 0}
	q3 := orderedmap.Pair[int, int]{Key: 60, Value: 0}
	r3 := orderedmap.Pair[int, int]{Key: 50, Value: 0}
	s3 := orderedmap.Pair[int, int]{Key: 40, Value: 0}
	t3 := orderedmap.Pair[int, int]{Key: 30, Value: 0}
	u3 := orderedmap.Pair[int, int]{Key: 20, Value: 0}
	v3 := orderedmap.Pair[int, int]{Key: 10, Value: 0}

	renkoMap.AddPairs(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z,
		a1, b1, c1, d1, e1, f1, g1, h1, i1, j1, k1, l1, m1, n1, o1, p1, q1, r1, s1, t1, u1, v1, w1, x1, y1, z1,
		a2, b2, c2, d2, e2, f2, g2, h2, i2, j2, k2, l2, m2, n2, o2, p2, q2, r2, s2, t2, u2, v2, w2, x2, y2, z2,
		a3, b3, c3, d3, e3, f3, g3, h3, i3, j3, k3, l3, m3, n3, o3, p3, q3, r3, s3, t3, u3, v3)

	return renkoMap
}
