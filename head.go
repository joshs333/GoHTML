package GoHTML

import (
	"anden-wvs/GoHTML/css"
	"strconv"
)

//HEAD struct code
type HEAD struct {
	css     *css.CSS
	title   string
	icon    string
	vWidth  int
	styles  []string
	scripts []string
}

func NewHEAD(css *css.CSS, title string, wid int) HEAD {
	return HEAD{css, title, "/img/j.png", wid, []string{}, []string{""}}
}

func (h HEAD) setCSS(p *css.CSS) {
	h.css = p
}

func (h HEAD) toHTML(path string) string {
	ret := ""
	ret += "<head>\n\t<title>" + h.title + "</title>\n"
	for j := len(h.styles); j > 0; j-- {
		ret += "\t<link rel=\"stylesheet\" type=\"text/css\" href=\"" + RelPath(path, h.styles[j]) + "\">\n"
	}

	ret += "\t<!-- [if lt IE 9] >\n\t\t<script src=\"http://html5shim.googlecode.com/svn/trunk/html5.js\"></script>\n\t<! [endif] -->\n"
	ret += "<script src=\"https://ajax.googleapis.com/ajax/libs/angularjs/1.6.4/angular.min.js\"></script>\n<script src=\"https://ajax.googleapis.com/ajax/libs/angularjs/1.6.4/angular-route.js\"></script>\n"

	if len(h.icon) > 0 {
		ret += "\t<link rel=\"apple-touch-icon-precomposed\" href=\"" + RelPath(path, h.icon) + "\">\n"
		ret += "\t<link rel=\"icon\" href=\"" + RelPath(path, h.icon) + "\">\n"
	}

	ret += "\t<meta http-equiv=\"Content-Type\" content=\"text/GoHTML; charset=UTF-8\">\n"
	ret += "\t<meta name=\"viewport\" content=\"width=" + strconv.Itoa(h.vWidth) + "px, initial-scale=1.0\">\n"
	ret += "<style>\n" + h.css.ToCSS() + "</style>"
	ret += "</head>\n"

	return ret
}
