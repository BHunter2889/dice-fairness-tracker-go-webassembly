package main

import (
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
	//s.populateGrid()
}

func newRollCounters(numSides int) []SideRollCounterData {
	ret := make([]SideRollCounterData, numSides)
	for i := 0; i < numSides; i++ {
		ret[i] = SideRollCounterData{
			SideNumber: i+1,
			RowIndex: i%4,
			IndexOfRow: int(i/4),
		}
	}
	return ret
}

func (s *globalState) populateGrid() {
	tempGrid := make([][]SideRollCounterData, s.NumSides)
	count := 0
	for i := 0; i < s.NumRows; i++ {
		for j := 0; j < 4 && count < s.NumSides; j++ {
			tempGrid[i][j] = s.SideRollCounters[count]
			count++
		}
	}
	println(tempGrid)
	s.CounterGrid = tempGrid
}