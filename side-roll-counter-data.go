package main

import "github.com/vugu/vugu"

type SideRollCounterData struct {
	SideNumber int
	RowIndex int
	IndexOfRow int
	Count int `default:"0"`
	StyleContentString string
}

func (comp *SideRollCounter) NewData(props vugu.Props) (interface{}, error) {
	sideNumber := props["side-number"].(int)
	ret := &state.SideRollCounters[sideNumber-1]
	return ret, nil
}

func (data *SideRollCounterData) Increment() {
	data.Count++
	IncrementStateTotal()
}

func (data *SideRollCounterData) Decrement() {
	if data.Count > 0 {
		data.Count--
		DecrementStateTotal()
	}
}
