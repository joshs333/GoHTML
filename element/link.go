package element

import "anden-wvs/GoHTML/css"


//URLType used for goJS
type URLType int

const (
	internal URLType = 0
	external URLType = 1
	relative URLType = 2
)

func (l URLType) String() string {
	switch l {
	case internal:
		return "int"
	case external:
		return "ext"
	case relative:
		return "rel"
	}
	return "int"
}

//Link data struct generates an GoHTML A element
type Link struct {
	attrs    map[string]string
	style    css.CSS
	contents []Element
}

//a nil link empty link objects
var NilLink = &Link{nil, css.NilCss, nil}

//NewLink() returns a pointer to a new Link object
func NewLink() *Link {
	return &Link{make(map[string]string), css.NewCSS(), []Element{}}
}

//Clone returns an exact clone of the given link object, is a deep copy
func (a *Link) Clone() Element {
	newAttrs := make(map[string]string)
	for k, v := range a.attrs {
		newAttrs[k] = v
	}
	var newContents []Element
	for i := 0; i < len(a.contents); i++ {
		newContents = append(newContents, a.contents[i].Clone())
	}
	return &Link{newAttrs, a.style.Clone(), newContents}
}


//GetType() returns the type of element Link is in string form ("a")
func (a *Link) GetType() string {
	return "a"
}

//ToHTML() returns the given link object in GoHTML form
func (a *Link) ToHTML() string {
	ret := "<" + a.GetType()

	//prints the element attributes
	for prop, val := range a.attrs {
		ret += " " + prop + "=\"" + val + "\""
	}

	//only prints style if it exists
	if !a.style.IsEmpty() {
		ret += " style=\"" + a.style.ToInline() + "\""
	}
	ret += ">"

	for i := 0; i < len(a.contents); i++ {
		ret += a.contents[i].ToHTML()
	}
	ret += "</" + a.GetType() + ">"
	return ret
}

//GetAttr() returns the value of the given attr for the given link
func (a *Link) GetAttr(attr string) string {
	if value, ok := a.attrs[attr]; ok {
		return value
	} else {
		return ""
	}
}


//SetAttr() sets the given attr to prop for the given Link
func (a *Link) SetAttr(attr string, prop string) {
	if attr != "style" && attr != "ID" {
		a.attrs[attr] = prop
	}
}

//GetId() returns the ID of the given link object
func (a *Link) GetId() string {
	return a.GetAttr("id")
}

//SetId() sets the ID of the given link object
func (a *Link) SetId(ID string) {
	a.SetAttr("id", ID)
}

//GetCSS() return the value of the given prop for the given Link
func (a *Link) GetCSS(prop string) string {
	return a.style.GetProp(prop)
}

//SetCSS() sets the given prop to val for the given Link
func (a *Link) SetCSS(prop string, val string) {
	a.style.SetProp(prop, val)
}

//SetText() sets the internal text to the given Link to text.
func (a *Link) SetText(text string) {
	for i := 0; i < len(a.contents); i++ {
		//finds the span element
		if a.contents[i].GetType() == "span" {
			//sets its text
			a.contents[i].SetText(text)
			return
		}
	}
	//if no such element exists, adds it to set text
	a.contents = append(a.contents, NewText(text))
}

//GetText() returns a reference to the Text element in teh given Link
func (a *Link) GetText() (*Text, bool) {
	for i := 0; i < len(a.contents); i++ {
		//finds the first span element
		if elem := a.contents[i]; elem.GetType() == "span" {
			return elem.(*Text), true
		}
	}
	//if none exists, returns nil and false
	return &NilText, false
}

//SetURLType() sets the url type of the given Link
func (a *Link) SetURLType(ltype URLType) {
	a.SetAttr("gnType", ltype.String())
}

//SetURL() sets the url of the given Link to href
func (a *Link) SetURL(href string) {
	a.SetAttr("href", href)
}
