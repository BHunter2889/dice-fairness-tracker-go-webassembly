package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"

type DiceSidesGridData struct {
	NumSides int `default:"20"`
	NumRows  int
	Counters []SideRollCounterData
}

func (comp *DiceSidesGrid) NewData(props vugu.Props) (interface{}, error) {
	numSides := props["sides"].(int)
	numRows := props["rows"].(int)

	ret := &DiceSidesGridData{
		NumSides: numSides,
		NumRows:  numRows,
		Counters: state.SideRollCounters,
	}
	return ret, nil
}

var _ vugu.ComponentType = (*DiceSidesGrid)(nil)

func (comp *DiceSidesGrid) BuildVDOM(dataI interface{}) (vdom *vugu.VGNode, css *vugu.VGNode, reterr error) {
	data := dataI.(*DiceSidesGridData)
	_ = data
	_ = fmt.Sprint
	_ = reflect.Value{}
	event := vugu.DOMEventStub
	_ = event
	css = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "style", DataAtom: vugu.VGAtom(458501), Namespace: "", Attr: []vugu.VGAttribute(nil)}
	css.AppendChild(&vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t.dice-sides-grid {\n\t\twidth: 100%;\n\t\tmax-height: available;\n\t\tmin-height: available;\n\t}\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)})
	var n *vugu.VGNode
	n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "dice-sides-grid"}}}
	vdom = n
	{
		parent := n
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		for i := 0; i < data.NumSides; i = i + 4 {
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "counter-grid-row", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n.Props = vugu.Props{
				"side-roll-counters": data.Counters[i : i+4],
			}
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
	}
	return
}

type DiceSidesGrid struct {}

func init() { vugu.RegisterComponentType("dice-sides-grid", &DiceSidesGrid{}) }
