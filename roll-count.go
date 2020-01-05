package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"

type RollCountData struct {
	Count int `default:"0"`
}

func (data *RollCountData) Increment() {
	data.Count++
	fmt.Println(data.Count)
}

var _ vugu.ComponentType = (*RollCount)(nil)

func (comp *RollCount) BuildVDOM(dataI interface{}) (vdom *vugu.VGNode, css *vugu.VGNode, reterr error) {
	data := dataI.(*RollCountData)
	_ = data
	_ = fmt.Sprint
	_ = reflect.Value{}
	event := vugu.DOMEventStub
	_ = event
	css = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "style", DataAtom: vugu.VGAtom(458501), Namespace: "", Attr: []vugu.VGAttribute(nil)}
	css.AppendChild(&vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t.dice-side-number-tracker {\n\t\tbackground: #eee;\n\t}\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)})
	var n *vugu.VGNode
	n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "dice-side-number-tracker"}}}
	vdom = n
	{
		parent := n
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", DataAtom: vugu.VGAtom(102662), Namespace: "", Attr: []vugu.VGAttribute(nil)}
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
		{
			parent := n
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Nat20", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n.InnerHTML = fmt.Sprint(data.Count)
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
	}
	return
}

type RollCount struct {}

func (ct *RollCount) NewData(props vugu.Props) (interface{}, error) { return &RollCountData{}, nil }

func init() { vugu.RegisterComponentType("roll-count", &RollCount{}) }
