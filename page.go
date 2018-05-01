/*
	page.go will contain structs, interfaces, and
	functions concerning full pages
*/
package GoHTML

import (
	"anden-wvs/GoHTML/body"
	"anden-wvs/GoHTML/css"
	"anden-wvs/GoHTML/js"
)

//PAGE struct code
type PAGE struct {
	style []*css.CSS
	path  string
	blocks []body.Block
	scripts    []js.Module
	containers map[string]Container
}

func NewPAGE(path string) PAGE {
	tcss := css.NewCSS()
	tcss.SetElem("body")
	ret := PAGE{[]*css.CSS{&tcss}, path, []body.Block{}, []js.Module{}, make(map[string]Container)}
	return ret
}

func (p *PAGE) Clone() PAGE {
	var newCSS []*css.CSS
	for i := 0; i < len(p.style); i++ {
		tcss := p.style[i].Clone()
		newCSS = append(newCSS, &tcss)
	}
	return *p
}

func (p *PAGE) Head(path string) string {
	ret := ""

	ret += "\t<!-- [if lt IE 9] >\n\t\t<script src=\"http://html5shim.googlecode.com/svn/trunk/html5.js\"></script>\n\t<! [endif] -->\n"
	ret += "<script src=\"https://ajax.googleapis.com/ajax/libs/angularjs/1.6.4/angular.min.js\"></script>\n<script src=\"https://ajax.googleapis.com/ajax/libs/angularjs/1.6.4/angular-route.js\"></script>\n"

	ret += "<script>\n"
	for i := 0; i < len(p.scripts); i++ {
		ret += p.scripts[i].ToJS()
	}
	ret += "</script>\n"

	ret += "\t<meta http-equiv=\"Content-Type\" content=\"text/GoHTML; charset=UTF-8\">\n"
	ret += "\t<meta name=\"viewport\" content=\"width=1280px, initial-scale=1.0\">\n"
	ret += "<style>\n"
	for i := 0; i < len(p.style); i++ {
		ret += p.style[i].ToCSS()
	}
	ret += "</style>\n"
	ret += "</head>\n"

	return ret
}

func (p *PAGE) ToHTML() string {
	ret := "<!DOCTYPE GoHTML>\n<GoHTML lang=\"en\">\n"
	ret += p.Head(p.path)
	ret += "<body>"
	for i := 0; i < len(p.blocks); i++ {
		ret += p.blocks[i].ToHTML()
	}
	ret += "</body>"
	ret += "</GoHTML>"
	return ret
}

func (p *PAGE) AddScript(script js.Module) {
	p.scripts = append(p.scripts, script)
}

func (p *PAGE) AddBlock(newBlock body.Block) {
	p.blocks = append(p.blocks, newBlock)
	p.style = append(p.style, newBlock.GetCSSAddr())
}

func (p *PAGE) GetBlock(index int) (body.Block, bool){
	if index < len(p.blocks) {
		return p.blocks[index], true
	}
	return nil, false
}

func (p *PAGE) AddContainer(path string, content Container) {
	p.containers[path] = content
}

func (p *PAGE) GetContainer(path string) (Container, bool) {
	if content, ok := p.containers[path]; ok {
		return content, ok
	}
	return NilContainer, false
}
