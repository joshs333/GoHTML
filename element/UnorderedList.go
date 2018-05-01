package element

import (
	"anden-wvs/GoHTML/css"
)

type UnorderedList struct {
	attrs    map[string]string
	style    css.CSS
	contents []*ListItem
}

func NewUnorderedList() *UnorderedList {
	return &UnorderedList{make(map[string]string), css.NewCSS(), []*ListItem{}}
}

func (u *UnorderedList) GetType() string {
	return "ul"
}

func (u *UnorderedList) ToHTML() string {
	ret := "<" + u.GetType()

	for prop, val := range u.attrs {
		ret += " " + prop + "=\"" + val + "\""
	}

	if !u.style.IsEmpty() {
		ret += " style=\"" + u.style.ToInline() + "\""
	}
	ret += ">\n"

	for i := 0; i < len(u.contents); i++ {
		ret += u.contents[i].ToHTML()
	}

	ret += "</" + u.GetType() + ">\n"
	return ret
}

func (u *UnorderedList) GetAttr(attr string) string {
	if value, ok := u.attrs[attr]; ok {
		return value
	} else {
		return ""
	}
}

func (u *UnorderedList) SetAttr(attr string, prop string) {
	if attr != "style" && attr != "ID" {
		u.attrs[attr] = prop
	}
}

func (u *UnorderedList) GetId() string {
	return u.GetAttr("id")
}

func (u *UnorderedList) SetId(ID string) {
	u.SetAttr("id", ID)
	u.style.SetElem(ID)
}

func (u *UnorderedList) GetCSS(prop string) string {
	return u.style.GetProp(prop)
}

func (u *UnorderedList) SetCSS(prop string, val string) {
	u.style.SetProp(prop, val)
}

//setText will set the text of the last list item
func (u *UnorderedList) SetText(text string) {
	if len(u.contents) > 0 {
		u.contents[len(u.contents)-1].SetText(text)
	}
	return
}

//getText will return a pointer to the text Element of the last list item
//or will return NilText, false if no such item exists
func (u *UnorderedList) GetText() (*Text, bool) {
	if len(u.contents) > 0 {
		return u.contents[len(u.contents)-1].GetText()
	} else {
		return &NilText, false
	}
}

//getElemNum returns the # of elems contained
func (u *UnorderedList) GetItemCount() int {
	return len(u.contents)
}

//getListItem returns the n + 1th item
//if n > # of elems contained, will return the last item
//if no elements contained, returns empty ListItem
func (u *UnorderedList) GetListItem(index int) *ListItem {
	//if index will exceed bounds of array
	if index >= len(u.contents) {
		if len(u.contents) == 0 { //if contains no items
			rt := NewListItem()
			return rt
		}
		return u.contents[len(u.contents)-1] //return last item
	}
	return u.contents[index] //returns item at index
}

//SetListItem will set the ListItem at index n to given ListItem
//if n > # of ListItem - 1, the given ListItem is amended to the end
func (u *UnorderedList) SetListItem(index int, element *ListItem) {
	if index >= len(u.contents) {
		u.AddListItem(element)
	}
	u.contents[index] = element
}

func (u *UnorderedList) AddListItem(element *ListItem) {
	u.contents = append(u.contents, element)
}

//getListItem removes the n + 1th item
//if no elements contained or n > # of items, no action will be done
func (u UnorderedList) RemoveListItem(i int) {
	if i >= len(u.contents) {
		return
	}
	//this code was taken from https://github.com/golang/go/wiki/SliceTricks
	//as a non-memory leak inducing method of deleting slice elements
	//thanks Wiki's! :)
	copy(u.contents[i:], u.contents[i+1:])
	u.contents[len(u.contents)-1] = NilListItem
	u.contents = u.contents[:len(u.contents)-1]
}
