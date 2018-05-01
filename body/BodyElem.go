package body

import "anden-wvs/GoHTML/css"

//import "anden-wvs/GoHTML/element"

/*
	Interface for main divs that will hold other elements in the body
*/
type Block interface {
	//should have func New<DIV>() <DIV>
	//should have var Nil<DIV>() <DIV>

	ToHTML() string  //to get HTML code for printing web page
	GetType() string //gets the div type to classify the different types
	//Clone() Block

	SetAttr(string, string)
	GetAttr(string) string

	GetId() string   //returns the div ID <-- used for Angular and CSS
	SetId(string)    //sets the DIV id <-- used for Angular and CSS

	//I will implement this in NAV shortly and will reapply it to the interface
	//AddElem(element.Element)     //adds and element to the div
	//GetElem(int)     //gets an element by order
	//GetElemById(string) //returns first element with ID of given string

	SetCSS(string, string)
	GetCSS(string) string
	GetCSSAddr() *css.CSS
}
