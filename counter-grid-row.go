package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"

type CounterGridRowData struct {
	sideRollCountersData []SideRollCounterData
}

func (comp *CounterGridRow) NewData(props vugu.Props) (interface{}, error) {
	ret := &CounterGridRowData{}
	ret.sideRollCountersData, _ = props["side-roll-counters"].([]SideRollCounterData)
	return ret, nil
}

var _ vugu.ComponentType = (*CounterGridRow)(nil)

func (comp *CounterGridRow) BuildVDOM(dataI interface{}) (vdom *vugu.VGNode, css *vugu.VGNode, reterr error) {
	data := dataI.(*CounterGridRowData)
	_ = data
	_ = fmt.Sprint
	_ = reflect.Value{}
	event := vugu.DOMEventStub
	_ = event
	css = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "style", DataAtom: vugu.VGAtom(458501), Namespace: "", Attr: []vugu.VGAttribute(nil)}
	css.AppendChild(&vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t.counter-grid-row {\n\t\twidth: 90%;\n\t\tdisplay: inline-flex;\n\t\tjustify-content: space-evenly;\n\t}\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)})
	var n *vugu.VGNode
	n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "span", DataAtom: vugu.VGAtom(40708), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "counter-grid-row"}}}
	vdom = n
	{
		parent := n
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		for _, counter := range data.sideRollCountersData {
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "side-roll-counter", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n.Props = vugu.Props{
				"side-number": counter.SideNumber,
			}
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
	}
	return
}

type CounterGridRow struct {}

func init() { vugu.RegisterComponentType("counter-grid-row", &CounterGridRow{}) }
