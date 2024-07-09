package rythm

import (
	// "fmt"
	"testing"

	"desktop.rythm/core/data"
)

func TestRealtimeChannel(t *testing.T) {
	var prices = make(chan data.Candle)
	go watchStream(prices)

	for value := range prices {
		// fmt.Printf("%v here is", value)
		if value.Direction == 3 {
			t.Fail()
		}
	}

}
