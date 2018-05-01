package body

import (
	"anden-wvs/GoHTML/css"
	"anden-wvs/GoHTML/element"
)

type Main struct {
	css      css.CSS
	attrs    map[string]string
	elements []element.Element
}

func NewMain() *Main {
	return &Main{css.NewCSS(), make(map[string]string), []element.Element{}}
}

func (m *Main) Clone() Block {
	newAttrs := make(map[string]string)
	for k, v := range m.attrs {
		newAttrs[k] = v
	}
	var NewElems []element.Element
	for i := 0; i < len(m.elements); i++ {
		NewElems = append(NewElems, m.elements[i].Clone())
	}
	return &Main{m.css.Clone(), newAttrs, NewElems}
}

func (m *Main) ToHTML() string {
	ret := "<" + m.GetType()
	for prop, val := range m.attrs {
		ret += " " + prop + "=\"" + val + "\" "
	}
	ret += ">"
	for i := 0; i < len(m.elements); i++ {
		ret += m.elements[i].ToHTML()
	}
	ret += "</" + m.GetType() + ">\n"
	return ret
}

func (m *Main) GetType() string {
	return "main"
}

func (m *Main) SetAttr(attr string, value string) {
	m.attrs[attr] = value
}

func (m *Main) GetAttr(attr string) string {
	if val, ok := m.attrs[attr]; ok {
		return val
	}
	return ""
}

func (m *Main) GetId() string {
	return m.GetAttr("id")
}

func (m *Main) SetId(id string) {
	m.SetAttr("id", id)
}

func (m *Main) AddElem(elem element.Element) {
	m.elements = append(m.elements, elem)
}

func (m *Main) GetElem(index int) (element.Element, bool) {
	if index < len(m.elements) {
		return m.elements[index], true
	}
	return nil, false
}

func (m *Main) GetElemById(id string) []element.Element {
	var ret []element.Element
	for i := 0; i < len(m.elements); i++ {
		if m.elements[i].GetId() == id {
			ret = append(ret, m.elements[i])
		}
	}
	return ret
}

func (m *Main) SetCSS(property string, value string) {
	m.css.SetProp(property, value)
}

func (m *Main) GetCSS(property string) string {
	return m.css.GetProp(property)
}

func (m *Main) GetCSSAddr() *css.CSS {
	return &m.css
}
