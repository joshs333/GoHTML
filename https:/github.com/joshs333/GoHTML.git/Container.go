package GoHTML

import (
	"anden-wvs/GoHTML/element"
)

type Container struct {
	elements []element.Element
}

func NewContainer() Container {
	return Container{[]element.Element{}}
}

var NilContainer = Container{nil}

func (m *Container) Clone() *Container {
	var NewElems []element.Element
	for i := 0; i < len(m.elements); i++ {
		NewElems = append(NewElems, m.elements[i].Clone())
	}
	return &Container{NewElems}
}

func (m *Container) ToHTML() string {
	ret := ""
	for i := 0; i < len(m.elements); i++ {
		ret += m.elements[i].ToHTML()
	}
	return ret
}

func (m *Container) GetType() string {
	return "Container"
}

func (m *Container) AddElem(elem element.Element) {
	m.elements = append(m.elements, elem)
}

func (m *Container) GetElem(index int) (element.Element, bool) {
	if index < len(m.elements) {
		return m.elements[index], true
	}
	return nil, false
}

func (m *Container) GetElemById(id string) []element.Element {
	var ret []element.Element
	for i := 0; i < len(m.elements); i++ {
		if m.elements[i].GetId() == id {
			ret = append(ret, m.elements[i])
		}
	}
	return ret
}
