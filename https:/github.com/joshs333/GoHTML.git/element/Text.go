package element

import "anden-wvs/GoHTML/css"

//Text acts as both the Text Element and also as just a text-holder in other elements that
//contain text
type Text struct {
	text  string
	style css.CSS
	attrs map[string]string
}

var NilText = Text{"", css.NilCss, nil}

//NewText initializes a Text with nil ID and no data so it will act as only a text-holder
//which is the original reason Text was coded into JHtml, to make simple text, but also make
//it stylizable with CSS
func NewText(text string) *Text {
	return &Text{text, css.NewCSS(), make(map[string]string)}
}

func (s *Text) Clone() Element {
	newAttrs := make(map[string]string)
	for k, v := range s.attrs {
		newAttrs[k] = v
	}
	return &Text{s.text, s.style.Clone(), newAttrs}
}

func (s *Text) GetType() string {
	return "span"
}

func (s *Text) ToHTML() string {
	//if the Text has no properties...
	if s.style.IsEmpty() && len(s.attrs) <= 0 {
		return s.text //it will act as just a text-holder...
	} else { //if it has properties, it will generate a span
		ret := "<span"

		for prop, val := range s.attrs {
			ret += " " + prop + "=\"" + val + "\" "
		}

		if !s.style.IsEmpty() {
			ret += " style=\"" + s.style.ToInline() + "\""
		}

		ret += ">" + s.text
		ret += "</span>"
		return ret
	}
}

func (s *Text) GetAttr(attr string) string {
	return s.attrs[attr]
}

func (s *Text) SetAttr(attr string, prop string) {
	if attr != "style" && attr != "ID" {
		s.attrs[attr] = prop
	}
}

func (s *Text) GetId() string {
	return s.GetAttr("id")
}

func (s *Text) SetId(ID string) {
	s.SetAttr("id", ID)
	s.style.SetElem(ID)
}

func (s *Text) GetCSS(prop string) string {
	return s.style.GetProp(prop)
}

func (s *Text) SetCSS(prop string, val string) {
	s.style.SetProp(prop, val)
}

func (s *Text) SetText(text string) {
	s.text = text
}

func (s *Text) GetText() (*Text, bool) {
	return s, true
}
