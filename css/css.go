package css

type CSS struct {
	elem       string
	properties map[string]string
	sub        map[string]*CSS
}

var NilCss = CSS{"", nil, nil}

func NewCSS() CSS {
	return CSS{"", make(map[string]string), nil}
}

func Style(styles []CSS) {
	ret := "<style>\\n"
	for i := 0; i < len(styles); i++ {
		ret := styles[i].elem + " {\n"
		for prop, val := range styles[i].properties {
			ret += "\t" + prop + ": " + val + ";\n"
		}
		ret += "}\n"
	}
	ret += "<\\style>\n"
}

func (c *CSS) Clone() CSS {
	newProp := make(map[string]string)
	for key, value := range c.properties {
		newProp[key] = value
	}
	if c.sub != nil {
		retSub := make(map[string]*CSS)
		for k := range c.sub {
			cloned := c.sub[k].Clone()
			retSub[k] = &cloned
		}
		return CSS{c.elem, newProp, retSub}
	} else {
		return CSS{c.elem, newProp, nil}
	}
}

func (c *CSS) SetProp(property string, value string) {
	c.properties[property] = value
}

func (c *CSS) GetProp(property string) string {
	return c.properties[property]
}

func (c *CSS) ToCSS() string {
	ret := c.elem + " {\n"
	for prop, val := range c.properties {
		ret += "\t" + prop + ": " + val + ";\n"
	}
	ret += "}\n"
	return ret
}

func (c *CSS) ToInline() string {
	ret := ""
	for prop, val := range c.properties {
		ret += prop + ": " + val + "; "
	}
	return ret
}

func (c *CSS) IsEmpty() bool {
	return !(len(c.properties) > 0)
}

func (c *CSS) SetElem(ID string) {
	c.elem = ID
}

func (c *CSS) SetChild(child string, property string, value string) {
	if c.sub == nil {
		c.sub = make(map[string]*CSS)
		new := NewCSS()
		c.sub["*"] = &new
	}
	if _, ok := c.sub[child]; !ok {
		new := NewCSS()
		c.sub[child] = &new
	}
	c.sub[child].SetProp(property, value)
}

func (c *CSS) SetAllChildren(property string, value string) {
	if c.sub == nil {
		c.sub = make(map[string]*CSS)
		new := NewCSS()
		c.sub["*"] = &new
	}
	c.sub["*"].SetProp(property, value)
}
