package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"

var _ vugu.ComponentType = (*SideRollCounter)(nil)

func (comp *SideRollCounter) BuildVDOM(dataI interface{}) (vdom *vugu.VGNode, css *vugu.VGNode, reterr error) {
	data := dataI.(*SideRollCounterData)
	_ = data
	_ = fmt.Sprint
	_ = reflect.Value{}
	event := vugu.DOMEventStub
	_ = event
	css = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "style", DataAtom: vugu.VGAtom(458501), Namespace: "", Attr: []vugu.VGAttribute(nil)}
	css.AppendChild(&vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t.side-roll-counter:hover {\n\t\t/*box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2);*/\n\t\tfilter: drop-shadow(0 8px 5px black);\n\t}\n\n\t/*@media (min-width: 455px) {*/\n\t\t/**/\n\t/*}*/\n\t.side-roll-counter {\n\t\t/*box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);*/\n\t\t/*transition: 0.3s;*/\n\t\t/*!*width: auto;*!*/\n\t\t/*margin: 20px;*/\n\t\t/*position: relative;*/\n\t\t/*width: 81px;*/\n\t\t/*height: 95px;*/\n\t\t/*background: 50% transparent url(https://www.dndbeyond.com/Content/Skins/Waterdeep/images/character-sheet/content-frames/abilityscore.svg) no-repeat;*/\n\t\t/*background-size: contain;*/\n\t\t/*text-align: center;*/\n\n\t\t/* TODO: Decide whether this is better: */\n\t\t/*box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2), inset 0 0 0 2000px rgba(255,255,255,0.1);*/\n\t\ttransition: 0.3s;\n\t\t/* width: auto; */\n\t\tmargin: 10px;\n\t\tposition: relative;\n\t\t/*width: 81px;*/\n\t\t/*height: 95px;*/\n\t\tz-index: 100;\n\t\tbackground: 50% transparent url(https://www.dndbeyond.com/Content/Skins/Waterdeep/images/character-sheet/content-frames/abilityscore.svg) no-repeat;\n\t\tbackground-size: contain;\n\t\ttext-align: center;\n\t\tfilter: drop-shadow(1px 2px 2px black);\n\t\tplace-self: stretch;\n\t\tplace-items: center;\n\t\tdisplay: grid;\n\t\tgrid-template-columns: 32px minmax(5px, 1fr) minmax(5px, 1fr) minmax(5px, 1fr) 32px;\n\t\tgrid-template-rows: 30% 1fr 30%;\n\t}\n\n\t.container {\n\t\tpadding: 2px 10%;\n\t\tdisplay: flex;\n\t\t/*height: 75px;*/\n\t\tjustify-content: space-around;\n\t}\n\n\t.roll-counter-count {\n\t\t/*position: absolute;*/\n\t\tfont-weight: 700;\n\t\t/*left: 0;*/\n\t\t/*right: 0;*/\n\t\t/*bottom: 0.3em;*/\n\t\tfont-size: xx-large;\n\t\tgrid-area: 3 / 3;\n\t}\n\n\tbutton {\n\t\twidth: 45px;\n\t\theight: 45px;\n\t\tmin-width: 45px;\n\t\tmin-height: 45px;\n\t\tfilter: drop-shadow(1px 1px 1px black);\n\t}\n\n\tbutton:hover {\n\t\tfilter: drop-shadow(2px 2px 2px black);\n\t}\n\n\tbutton:active {\n\t\tfilter: drop-shadow(0 0 1px black);\n\t\ttransform: translateY(1px);\n\t}\n\n\t/*.count-decrement:hover {*/\n\t\t/*filter: drop-shadow(2px 2px 2px black);*/\n\t/*}*/\n\n\t.count-decrement {\n\t\tbackground-size: cover;\n\t\tposition: relative;\n\t\tdisplay: inline-block;\n\t\tbackground-image: url(\"static/icons/minus.svg\");\n\t\tbackground-repeat: no-repeat;\n\t\tbackground-position: center;\n\t\tbackground-color: transparent;\n\t\tcolor: transparent;\n\t\tborder: 0;\n\t\tgrid-area: 2 / 2;\n\t\t/*width: 25px;*/\n\t\t/*height: 25px;*/\n\t\t/*float: left;*/\n\t\t/*margin: 0 6px 0 0;*/\n\t}\n\n\t/*.count-decrement:after {*/\n\t\t/*left:0;*/\n\t\t/*border-bottom:17.7px transparent;*/\n\t\t/*border-left:17.7px transparent;*/\n\t\t/*transform:rotate(45deg) skew(19deg,19deg);*/\n\t/*}*/\n\n\t/*.count-decrement:hover, .count-increment:hover {*/\n\t\t/*box-shadow: 0 0 16px  rgba(0,0,0,0.24), 0 0 50px rgba(0,0,0,0.19);*/\n\t/*}*/\n\n\t.count-increment {\n\t\tbackground-size: cover;\n\t\tposition: relative;\n\t\tdisplay: inline-block;\n\t\tbackground-image: url(\"static/icons/plus.svg\");\n\t\tbackground-repeat: no-repeat;\n\t\tbackground-position: center;\n\t\tbackground-color: transparent;\n\t\tcolor: transparent;\n\t\tborder: 0;\n\t\tgrid-area: 2 / 4;\n\t\t/*width: 25px;*/\n\t\t/*height: 25px;*/\n\t\t/*float: right;*/\n\t\t/*margin: 0 0 0 6px;*/\n\t}\n\n\t/*TODO: Why doesn't this (with the button rule using it) not work??? */\n\t@keyframes ripple {\n\t\t0% {\n\t\t\ttransform: scale(0, 0);\n\t\t\topacity: 1;\n\t\t}\n\t\t20% {\n\t\t\ttransform: scale(25, 25);\n\t\t\topacity: 1;\n\t\t}\n\t\t100% {\n\t\t\topacity: 0;\n\t\t\ttransform: scale(40, 40);\n\t\t}\n\t}\n\n\tbutton:focus:not(:active)::after {\n\t\tanimation: ripple 1s ease-out;\n\t}\n\n\t[class|=\"pressed\"] {\n\t\t/*background: #333 url(//stephy.mccdgm.net/images/patterns/gray_sand.png);*/\n\t\ttext-align: center;\n\t\tfont-size: xx-large;\n\t\tgrid-area: 1 / 3;\n\t}\n\n\t[class|=\"pressed\"] span {\n\t\tcolor: transparent;\n\t\tfont-family: 'Oswald', sans-serif;\n\t\tposition: relative;\n\t\tz-index: 0;\n\t}\n\n\t[class|=\"pressed\"] span:before, [class|=\"pressed\"] span:after {\n\t\tposition: absolute;\n\t\tleft: 0;\n\t\tcolor: #222; /* Fallback for non-webkit */\n\t}\n\n\t[class|=\"pressed\"] span:before {\n\t\tz-index: 1;\n\t\tbackground: -webkit-linear-gradient(transparent, transparent),\n\t\turl(//stephy.mccdgm.net/images/patterns/dark_wall.png);\n\t\tbackground: -o-linear-gradient(transparent, transparent);\n\t\t-webkit-background-clip: text;\n\t\t-webkit-text-fill-color: transparent;\n\t}\n\n\t[class|=\"pressed\"] span:after {\n\t\ttext-shadow: 0px 1px 2px rgba(255,255,255,0.2);\n\t\tz-index: -1;\n\t}\n\n\t.hexagon {\n\t\twidth: 25px;\n\t\theight: 25px;\n\t\tposition: relative;\n\t\tbackground-position: center;\n\t}\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)})
	var n *vugu.VGNode
	n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "side-roll-counter"}}}
	vdom = n
	{
		parent := n
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "style", DataAtom: vugu.VGAtom(458501), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n.InnerHTML = fmt.Sprint(data.StyleContentString)
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n.Props = vugu.Props{
			"class": fmt.Sprintf("pressed-%v", data.SideNumber),
		}
		{
			parent := n
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "span", DataAtom: vugu.VGAtom(40708), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n.InnerHTML = fmt.Sprint(data.SideNumber)
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", DataAtom: vugu.VGAtom(102662), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "count-decrement"}}}
		parent.AppendChild(n)
		// @click = { data.Decrement() }
		{
			var i_ interface{} = data
			idat_ := reflect.ValueOf(&i_).Elem().InterfaceData()
			var i2_ interface{} = data.Decrement
			i2dat_ := reflect.ValueOf(&i2_).Elem().InterfaceData()
			n.SetDOMEventHandler("click", vugu.DOMEventHandler{
				ReceiverAndMethodHash: uint64(idat_[0]) ^ uint64(idat_[1]) ^ uint64(i2dat_[0]) ^ uint64(i2dat_[1]),
				Method:                reflect.ValueOf(data).MethodByName("Decrement"),
				Args:                  []interface{}{},
			})
		}
		if false {
			// force compiler to check arguments for type safety
			data.Decrement()
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", DataAtom: vugu.VGAtom(102662), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "count-increment"}}}
		parent.AppendChild(n)
		// @click = { data.Increment() }
		{
			var i_ interface{} = data
			idat_ := reflect.ValueOf(&i_).Elem().InterfaceData()
			var i2_ interface{} = data.Increment
			i2dat_ := reflect.ValueOf(&i2_).Elem().InterfaceData()
			n.SetDOMEventHandler("click", vugu.DOMEventHandler{
				ReceiverAndMethodHash: uint64(idat_[0]) ^ uint64(idat_[1]) ^ uint64(i2dat_[0]) ^ uint64(i2dat_[1]),
				Method:                reflect.ValueOf(data).MethodByName("Increment"),
				Args:                  []interface{}{},
			})
		}
		if false {
			// force compiler to check arguments for type safety
			data.Increment()
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "roll-counter-count"}, vugu.VGAttribute{Namespace: "", Key: "type", Val: "text"}}}
		parent.AppendChild(n)
		n.InnerHTML = fmt.Sprint(*data.Count)
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
	}
	return
}

type SideRollCounter struct {}

func init() { vugu.RegisterComponentType("side-roll-counter", &SideRollCounter{}) }
