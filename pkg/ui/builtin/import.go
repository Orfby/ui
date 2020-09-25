package builtin

import (
	"encoding/xml"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/orfby/ui/pkg/ui/element"
)

//Type for an element that imports another design
type Import struct {
	//The import node is an element
	element.Element
	//It is also (technically) a layout
	element.Layout

	//The path to the design
	Path string `uixml:"http://github.com/orfby/ui/api/schema path"`
}

//Function to create a new import element
func NewImport(name xml.Name, parent element.IsLayout) element.IsElement {
	return &Import{Element: element.NewElement(name, parent)}
}

//The XML name of the import element
var ImportTypeName = xml.Name{Space: "http://github.com/orfby/ui/api/schema", Local:"Import"}

//Function to unmarshal an XML element into
//an element. This function is usually only
//called by xml.Unmarshal
func (e *Import) UnmarshalXML(d* xml.Decoder, start xml.StartElement) (err error) {
	//Unmarshal the element part of the layout
	err = e.Element.UnmarshalXML(d, start)
	if err != nil {return err}

	//Set the element's attributes
	err = element.SetAttrs(e, start.Attr)
	if err != nil {return err}

	//Create the root
	root, err := element.NewRoot(e, e.Path)
	if err != nil {return err}

	//Set it as the only child
	e.Children = []element.IsElement{root.IsElement}

	return d.Skip()
}

//Function to reset the element
func (e *Import) Reset() {
	e.Element.Reset()
	e.Layout.Reset()
}

//Function to determine whether
//the element is initialised
func (e *Import) IsInitialised() bool {
	return e.Element.IsInitialised() &&
		element.ChildrenAreInitialised(e)
}

//Function to initialise an element's
//position, width and height. Because
//it doesn't know the element's actual
//size, it won't set the width or height
//if the relative width or height is
//"match_content"
func (e *Import) Init(window *pixelgl.Window, bounds *pixel.Rect) (err error) {
	//Initialise the element part of the import
	err = e.Element.Init(window, bounds)
	if err != nil {return err}

	//Initialise the child
	err = e.GetChild(0).Init(window, bounds)
	if err != nil {return err}

	//If the width is meant to match the content size
	if e.GetRelWidth().MatchContent {
		//Set the width as the child's
		e.SetActualWidth(e.GetChild(0).GetActualWidth())
	}

	//If the height is meant to match the content size
	if e.GetRelHeight().MatchContent {
		//Set the height as the child's
		e.SetActualHeight(e.GetChild(0).GetActualHeight())
	}

	return nil
}

//Function that is called when there
//is a new event
func (e *Import) NewEvent(window *pixelgl.Window) {
	e.Element.NewEvent(window)
	e.Layout.NewEvent(window)
}

//Function to draw the element
func (e *Import) Draw() {
	//Draw the element
	e.Element.Draw()
	//Draw the layout
	element.DrawLayout(e)
}