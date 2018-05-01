package element

import (
	"anden-wvs/GoHTML/css"
)

type ListItem struct {
	attrs    map[string]string
	style    css.CSS
	contents []Element
}

var NilListItem = &ListItem{nil, css.NilCss, nil}

func NewListItem() *ListItem {
	return &ListItem{make(map[string]string), css.NewCSS(), []Element{}}
}

func (l *ListItem) Clone() Element {
	newAttrs := make(map[string]string)
	for k, v := range l.attrs {
		newAttrs[k] = v
	}
	var newContents []Element
	for i:= 0; i < len(l.contents); i++ {
		newContents = append(newContents, l.contents[i].Clone())
	}
	return &ListItem{newAttrs, l.style.Clone(), newContents}
}

func (l *ListItem) GetType() string {
	return "li"
}

func (l *ListItem) ToHTML() string {
	ret := "<" + l.GetType()

	for prop, val := range l.attrs {
		ret += " " + prop + "=\"" + val + "\""
	}

	if !l.style.IsEmpty() {
		ret += " style=\"" + l.style.ToInline() + "\""
	}
	ret += ">"

	for i := 0; i < len(l.contents); i++ {
		ret += l.contents[i].ToHTML()
	}

	ret += "</" + l.GetType() + ">\n"
	return ret
}

func (l *ListItem) GetAttr(attr string) string {
	if value, ok := l.attrs[attr]; ok {
		return value
	} else {
		return ""
	}
}

func (l *ListItem) SetAttr(attr string, prop string) {
	if attr != "style" && attr != "ID" {
		l.attrs[attr] = prop
	}
}

func (l *ListItem) GetId() string {
	return l.GetAttr("id")
}

func (l *ListItem) SetId(ID string) {
	l.SetAttr("id", ID)
	l.style.SetElem(ID)
}

func (l *ListItem) GetCSS(prop string) string {
	return l.style.GetProp(prop)
}

func (l *ListItem) SetCSS(prop string, val string) {
	l.style.SetProp(prop, val)
}

func (l *ListItem) SetText(text string) {
	for i := 0; i < len(l.contents); i++ {
		if l.contents[i].GetType() == "span" {
			l.contents[i].SetText(text)
			return
		}
	}

	l.AddElem(NewText(text))
}

func (l *ListItem) GetText() (*Text, bool) {
	for i := 0; i < len(l.contents); i++ {
		if elem := l.contents[i]; elem.GetType() == "span" {
			return elem.(*Text), true
		}
	}
	return &NilText, false
}

//AddElem adds given Element to the Element being performed on
func (l *ListItem) AddElem(elem Element) {
	l.contents = append(l.contents, elem)
}

//SetElem sets the n-1th Element to the given Element, or adds it to the end if not enough space
func (l *ListItem) SetElem(n int, elem Element) {
	leng := len(l.contents)
	if n > leng-1 {
		l.AddElem(elem)
	} else {
		l.contents[n] = elem
	}
}

//getElemNum returns the # of elems contained
func (l *ListItem) GetElemNum() int {
	return len(l.contents)
}

//getElem returns the Element at n - 1 Element in this Element
//if n > # of elems contained, will return the text Element (Text)
//if no span contained, will return the last Element
func (l *ListItem) GetElem(index int) Element {
	//if index will exceed bounds of array
	if index >= len(l.contents) {
		//counter starts at bottom, looks for "Text" to stop, or stops at end
		for index = 0; index < len(l.contents) && l.contents[index].GetType() != "span"; index++ {
		}
	}
	//returns requested Element, "Text" Element, or last Element.. in that preference depending on existence
	return l.contents[index]
}
