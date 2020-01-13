package main

import (
	"fmt"
	"math"
)

var state *globalState

type globalState struct {
	NumSides                    int `default:"20"`
	NumRows                     int `default:"5"`
	SideRollCounts              []int
	SideRollCounters            []SideRollCounterData   // Will need set after initState() is called and numSides is known.
	CounterGrid                 [][]SideRollCounterData // TODO: Is this needed? Will it help with Styling the counter grid?
	TotalRollCount              int `default:"0"`
	IsMinNumRollsMet			bool
	DieBalanceComputationValues *ComputedPearsonsChiSqValues
}

func initState() {
	state = &globalState{}
}

func (s *globalState) initGrid(numSides int) {
	s.NumSides = numSides
	s.NumRows = int(math.Ceil(float64(numSides / 4)))
	s.SideRollCounters, s.SideRollCounts = newRollCounters(numSides)
	s.DieBalanceComputationValues = NewComputedPCSValues(numSides)
	s.TotalRollCount = 0
}

func (s *globalState) GetRollCounterDataForSide(sideNumber int) *SideRollCounterData {
	return &s.SideRollCounters[sideNumber-1]
}

func newRollCounters(numSides int) (srcData []SideRollCounterData, counts []int) {
	srcData = make([]SideRollCounterData, numSides)
	counts = make([]int, numSides)
	for i := 0; i < numSides; i++ {
		counts[i] = 0
		srcData[i] = SideRollCounterData{
			SideNumber:         i + 1,
			RowIndex:           i % 4,
			IndexOfRow:         int(i / 4),
			StyleContentString: fmt.Sprintf(`.pressed-%v span:before, .pressed-%v span:after {content:"%v"}`, i+1, i+1, i+1),
			Count:              &counts[i],
		}
	}
	return
}

func IncrementRollCountsInState(sideNumber int) {
	state.SideRollCounts[sideNumber-1]++
	state.TotalRollCount++
	state.IsMinNumRollsMet = state.DieBalanceComputationValues.DieConstants.MinNumberOfRolls <= state.TotalRollCount
	if state.IsMinNumRollsMet {
		state.DieBalanceComputationValues.ComputePChSqValues(state.TotalRollCount, state.SideRollCounts)
	}
}

func DecrementRollCountsInState(sideNumber int) {
	if state.TotalRollCount > 0 && state.SideRollCounts[sideNumber-1] > 0 {
		state.SideRollCounts[sideNumber-1]--
		state.TotalRollCount--
		if state.DieBalanceComputationValues.DieConstants.MinNumberOfRolls <= state.TotalRollCount {
			state.DieBalanceComputationValues.ComputePChSqValues(state.TotalRollCount, state.SideRollCounts)
		}
	}
}