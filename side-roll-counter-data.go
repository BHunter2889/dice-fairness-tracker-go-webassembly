package main

import "github.com/vugu/vugu"

type SideRollCounterData struct {
	SideNumber int
	RowIndex int
	IndexOfRow int
	Count int `default:"0"`
}

func (comp *SideRollCounter) NewData(props vugu.Props) (interface{}, error) {
	sideNumber := props["side-number"].(int)
	ret := &state.SideRollCounters[sideNumber-1]
	return ret, nil
}

func (data *SideRollCounterData) Increment() {
	data.Count++
}
