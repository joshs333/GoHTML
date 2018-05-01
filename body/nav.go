package body

import (
	"anden-wvs/GoHTML/css"
	"anden-wvs/GoHTML/element"
)

type Nav struct {
	css   css.CSS
	ref   map[string]*element.ListItem
	attrs map[string]string
	list  *element.UnorderedList
}

func NewNAV() *Nav {
	new := Nav{css.NewCSS(), make(map[string]*element.ListItem), make(map[string]string), element.NewUnorderedList()}
	new.css.SetElem("nav")
	new.list.SetCSS("list-style-type", "none")
	return &new
}

func (n *Nav) ToHTML() string {
	ret := "<" + n.GetType()
	for prop, val := range n.attrs {
		ret += " " + prop + "=\"" + val + "\""
	}
	if !n.css.IsEmpty() {
		ret += " style=\"" + n.css.ToInline() + "\""
	}
	ret += ">\n"
	ret += n.list.ToHTML()
	ret += "</" + n.GetType() + ">\n"
	return ret
}

func (n *Nav) GetType() string {
	return "nav"
}

func (n *Nav) GetId() string {
	return n.GetAttr("id")
}

func (n *Nav) SetId(id string) {
	n.SetAttr("id", id)
	n.css.SetElem("#" + id)
}

func (n *Nav) SetCSS(property string, value string) {
	n.css.SetProp(property, value)
}

func (n *Nav) GetCSS(property string) string {
	return n.css.GetProp(property)
}

func (n *Nav) GetCSSAddr() *css.CSS {
	return &n.css
}

func (n *Nav) SetAttr(attr string, value string) {
	n.attrs[attr] = value
}

func (n *Nav) GetAttr(attr string) string {
	if val, ok := n.attrs[attr]; ok {
		return val
	}
	return ""
}

func (n *Nav) AddLink(path string, NavTitle string) *element.ListItem {
	newA := element.NewLink()
	newLI := element.NewListItem()

	newA.SetText(NavTitle)
	newA.SetURL(path)
	newA.SetAttr("gnLink", "true")
	newLI.AddElem(newA)

	n.list.AddListItem(newLI)
	n.ref[path] = n.list.GetListItem(n.list.GetItemCount() - 1)
	return n.ref[path]
}

func (n *Nav) RemoveLink(path string) {
	for i := 0; i < n.list.GetItemCount(); i++ {
		if n.list.GetListItem(i) == n.ref[path] {
			n.list.RemoveListItem(i)
		}
	}
}

func (n *Nav) GetLinkLI(path string) (*element.ListItem, bool) {
	value, ok := n.ref[path]
	return value, ok
}

func (n *Nav) GetLink(path string) (*element.Link, bool) {
	if value, ok := n.ref[path]; ok {
		for i := 0; i < value.GetElemNum(); i++ {
			if elem := value.GetElem(i); elem.GetAttr("gnLink") == "true" {
				return elem.(*element.Link), ok
			}
		}
	}
	return element.NilLink, false
}

func (n *Nav) GetList() *map[string]*element.ListItem {
	return &n.ref
}

func (n *Nav) HasLink(path string) bool {
	_, ok := n.ref[path]
	return ok
}
