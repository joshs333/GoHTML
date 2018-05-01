package element

type Element interface {
	//returns the type of Element it is "H1", "P", "Text",etc
	GetType() string
	//returns the Element in HTML form (similar to .toString() )
	ToHTML() string

	//getAttr returns the property of an attribute of an Element
	GetAttr(attr string) string
	//setAttr will set an attribute of the Element
	SetAttr(attr string, prop string)
	//returns the ID of the Element
	GetId() string
	//sets the ID of the Element
	SetId(ID string)

	//getCSS will return the value of a CSS property
	GetCSS(prop string) string
	//setCSS will set the value of a CSS property
	SetCSS(prop string, val string)

	//setText will set the text of the first "Text" Element contained. If none exists it will make one.
	SetText(text string)
	GetText() (*Text, bool)

	//Clone returns a deep clone of the given Element
	Clone() Element
}
