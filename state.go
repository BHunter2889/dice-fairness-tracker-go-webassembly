package main

import (
	"fmt"
	"log"
	"math"
)

var state *globalState

type globalState struct {
	NumSides int `default:"20"`
	NumRows int `default:"5"`
	SideRollCounters []SideRollCounterData // Will need set after initState() is called and numSides is known.
	CounterGrid [][]SideRollCounterData
	TotalRollCount int `default:"0"`
}

func initState() {
	state = &globalState{}
}

func (s *globalState) initGrid (numSides int) {
	s.NumSides = numSides
	s.NumRows = int(math.Ceil(float64(numSides/4)))
	s.SideRollCounters = newRollCounters(numSides)
}

func newRollCounters(numSides int) []SideRollCounterData {
	ret := make([]SideRollCounterData, numSides)
	for i := 0; i < numSides; i++ {
		ret[i] = SideRollCounterData{
			SideNumber: i+1,
			RowIndex: i%4,
			IndexOfRow: int(i/4),
			StyleContentString: fmt.Sprintf(`.pressed-%v span:before, .pressed-%v span:after {content:"%v"}`, i+1, i+1, i+1),
		}
		log.Printf("Style String %v: %s", i, ret[i].StyleContentString)
	}

	return ret
}

func IncrementStateTotal() {
	state.TotalRollCount++
}

func DecrementStateTotal() {
	state.TotalRollCount--
}