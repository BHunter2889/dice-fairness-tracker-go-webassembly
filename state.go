package main

import (
	"fmt"
	"math"
)

var state *globalState

const d20 = 20
const d20ChiTableLookup = 30.143

type globalState struct {
	NumSides int `default:"20"`
	NumRows int `default:"5"`
	SideRollCounts []int
	SideRollCounters []SideRollCounterData // Will need set after initState() is called and numSides is known.
	CounterGrid [][]SideRollCounterData
	TotalRollCount int `default:"0"`
	DieBalanceComputationValues ComputedPearsonsChiSqValues
}

type PearsonsChiSqOption struct {
	SideRollCount int `default:"0"`
	ExpectedRollCount float64 `default:"0.0"`
	Error int
	SquaredError int
}

type ComputedPearsonsChiSqValues struct {
	OptionComputations []PearsonsChiSqOption
	BalanceThreshold float64
	SumSquaredError int
	IsBalanced bool
}

func initState() {
	state = &globalState{}
}

func (s *globalState) initGrid (numSides int) {
	s.NumSides = numSides
	s.NumRows = int(math.Ceil(float64(numSides/4)))
	s.SideRollCounts = newRollCounts(numSides)

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
	}

	return ret
}

func newRollCounts(numSides int) []int {
	ret := make([]int, numSides)
	for i := 0; i < numSides; i++ {
		ret[i] = 0
	}

	return ret
}


func IncrementStateTotal(sideNumber int) {
	state.SideRollCounters[sideNumber-1].Count++
	state.TotalRollCount++
	computeBalanceThreshold()
}

func DecrementStateTotal(sideNumber int) {
	if state.TotalRollCount > 0  && state.SideRollCounters[sideNumber-1].Count > 0 {
		state.SideRollCounters[sideNumber-1].Count--
		state.TotalRollCount--
		computeBalanceThreshold()
	}
}

func computeBalanceThreshold() {
	state.DieBalanceComputationValues.BalanceThreshold = (float64(state.TotalRollCount)/float64(20))*d20ChiTableLookup
}