package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"

type RootData struct {
	NumSides             int  `default:"20"`
	NumRows              int  `default:"5"`
	TotalRollCount       *int `default:"0"`
	MinNumberOfRolls     int
	IsMinNumRollsMet     bool    `default:"false"`
	BalanceThreshold     float64 `default:"0.0"`   // Threshold for the SumSquaredError.
	ExpectedRollsPerSide float64 `default:"0.0"`   // `TotalRollCount/NumberOfSides`
	SumSquaredError      float64 `default:"0.0"`   // Total of all options SquaredError different from expected roll count. aka 'SSE'
	IsBalanced           bool    `default:"false"` // SSE <= BalanceThreshold
}

func (comp *Root) NewData(props vugu.Props) (interface{}, error) {

	ret := &RootData{
		//NumSides: state
		//NumRows: props["NumRows"].(int),
		TotalRollCount: &state.TotalRollCount,
		//MinNumberOfRolls: props["MinNumberOfRolls"].(int),
		//IsMinNumRollsMet: props["IsMinNumRollsMet"].(bool),
		//IsMinNumRollsMet:isMinRollsMet,
		//BalanceThreshold: props["BalanceThreshold"].(float64),
		//ExpectedRollsPerSide: props["ExpectedRollsPerSide"].(float64),
		//SumSquaredError: props["SumSquaredError"].(float64),
		//IsBalanced: props["IsBalanced"].(bool),
	}
	return ret, nil
}

var _ vugu.ComponentType = (*Root)(nil)

func (comp *Root) BuildVDOM(dataI interface{}) (vdom *vugu.VGNode, css *vugu.VGNode, reterr error) {
	data := dataI.(*RootData)
	_ = data
	_ = fmt.Sprint
	_ = reflect.Value{}
	event := vugu.DOMEventStub
	_ = event
	css = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "style", DataAtom: vugu.VGAtom(458501), Namespace: "", Attr: []vugu.VGAttribute(nil)}
	css.AppendChild(&vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\tdiv.root {\n\t\t/*background: rgb(32,56,42) radial-gradient(circle, rgba(32,56,42,0.4) 1%, rgba(33,31,31,0.5) 100%);*/\n\t\t/*background: inherit;*/\n\t\twidth: 100%;\n\t\theight: 100%;\n\t\t/*position: relative;*/\n\t\toverflow: hidden;\n\t\tbackground-size: cover;\n\t\t/*margin: 50px;*/\n\t\t/*max-width: 100%;*/\n\t\t/*max-height: 100%;*/\n\t\tposition: absolute;\n\n\t\t/* TODO: Maybe? */\n\t\t /*display: flex;*/\n\t\t /*justify-content: center;*/\n\t\t /*align-items: center;*/\n\t\t /*flex-flow: column;*/\n\t\tdisplay: grid;\n\t\tgrid-template-rows: 80% 1fr;\n\t\t/*grid-template-rows: [counter-grid-start] 80% [counter-grid-end stat-row-start] 1fr [stat-row-end];*/\n\t}\n\n\tdiv.root:before {\n\t\tcontent: close-quote;\n\t\tbackground: inherit;\n\t\tposition: absolute;\n\t\tleft: -10px;\n\t\tright: -10px;\n\t\ttop: -10px;\n\t\tbottom: -10px;\n\t\tbox-shadow: inset 0 0 0 2000px rgba(255,255,255,0.1);\n\t\tfilter: blur(1px);\n\t\t/*max-width: 100%;*/\n\t\t/*max-height: 100%;*/\n\t}\n\n\t.stat-score-block {\n\t\tposition: relative;\n\t\twidth: 81px;\n\t\theight: 95px;\n\t\tbackground: 50% transparent url(https://www.dndbeyond.com/Content/Skins/Waterdeep/images/character-sheet/content-frames/ac.svg) no-repeat;\n\t\tbackground-size: contain;\n\t\ttext-align: center;\n\t}\n\n\t.stat-score-block-row {\n\t\twidth: 100%;\n\t\t/*height: 20%;*/\n\t\tdisplay: flex;\n\t\tjustify-content: space-around;\n\t\tposition: absolute;\n\t\tgrid-row-start: 2;\n\t\tgrid-row-end: 3;\n\t}\n\n\n\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)})
	var n *vugu.VGNode
	n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "root"}}}
	vdom = n
	{
		parent := n
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "dice-sides-grid", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n.Props = vugu.Props{
			"rows":  state.NumRows,
			"sides": state.NumSides,
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "stat-score-block-row"}}}
		parent.AppendChild(n)
		{
			parent := n
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "stat-score-block"}}}
			parent.AppendChild(n)
			n.InnerHTML = fmt.Sprint(fmt.Sprint(*data.TotalRollCount, " | "))
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "stat-score-block"}}}
			parent.AppendChild(n)
			n.InnerHTML = fmt.Sprint(fmt.Sprint(state.DieBalanceComputationValues.DieConstants.MinNumberOfRolls, " | "))
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			if state.IsMinNumRollsMet {
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "stat-score-block"}}}
				parent.AppendChild(n)
				n.InnerHTML = fmt.Sprint(fmt.Sprint(state.DieBalanceComputationValues.BalanceThreshold, " | "))
			}
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			if state.IsMinNumRollsMet {
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "stat-score-block"}}}
				parent.AppendChild(n)
				n.InnerHTML = fmt.Sprint(fmt.Sprint(state.DieBalanceComputationValues.ExpectedRollsPerSide, " | "))
			}
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			if state.IsMinNumRollsMet {
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "stat-score-block"}}}
				parent.AppendChild(n)
				n.InnerHTML = fmt.Sprint(fmt.Sprint(state.DieBalanceComputationValues.SumSquaredError, " | "))
			}
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			if state.IsMinNumRollsMet {
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "stat-score-block"}}}
				parent.AppendChild(n)
				n.InnerHTML = fmt.Sprint(fmt.Sprint(state.DieBalanceComputationValues.IsBalanced))
			}
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		for i := 0; i < 15; i++ {
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "firefly"}}}
			parent.AppendChild(n)
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
	}
	return
}

type Root struct {}

func init() { vugu.RegisterComponentType("root", &Root{}) }
