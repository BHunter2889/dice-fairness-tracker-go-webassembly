package main

import "github.com/vugu/vugu"

type SideRollCounterData struct {
	SideNumber         int
	RowIndex           int    // TODO: Is this needed? Will it help with Styling the counter grid?
	IndexOfRow         int    //TODO: Is this needed? Will it help with Styling the counter grid?
	Count              *int   `default:"0"` // Only Needed for Display, arithmetic to be done on backing pointed value.
	StyleContentString string // Feels like a hack to get around the limitations of Vugu's present lack of dynamic CSS support
}

func (comp *SideRollCounter) NewData(props vugu.Props) (interface{}, error) {
	sideNumber := props["side-number"].(int)
	ret := state.GetRollCounterDataForSide(sideNumber)
	return ret, nil
}

func (data *SideRollCounterData) Increment() {
	IncrementRollCountsInState(data.SideNumber)
}

func (data *SideRollCounterData) Decrement() {
	DecrementRollCountsInState(data.SideNumber)
}
