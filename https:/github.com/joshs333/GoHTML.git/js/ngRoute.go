package js

import (
	"anden-wvs/GoHTML/body"
)

type NgRoute struct {
	nav *body.Nav
	block *body.Main
}

func NewNgRoute(nav *body.Nav, block *body.Main) *NgRoute {
	newModule := NgRoute{nav, block}
	newModule.block.SetAttr("ng-app", "gnMain")
	newModule.block.SetAttr("ng-view", "")
	return &newModule
}

func (n *NgRoute) Load() {

}

func (n *NgRoute) ToJS() string {
	ret := "var app = angular.module(\"gnMain\", [\"ngRoute\"]);	app.config(function($routeProvider) {	$routeProvider"
	for k := range *n.nav.GetList() {
		if elem, _ := n.nav.GetLink(k); elem.GetAttr("gnType") != "ext" {
			if elem.GetAttr("gnTouch") == "true" {
				k := k[0:]
				ret += ".when(\"" + k + ".gnLink" + "\", { templateUrl : \"" + k[0:len(k)-0] + "\" })"
				elem.SetURL("#!" + k + ".gnLink")
			} else {
				ret += ".when(\"" + k + ".gnLink" + "\", { templateUrl : \"" + k + "\" })"
				elem.SetURL("#!" + k + ".gnLink")
				elem.SetAttr("gnTouch", "true")
			}
		}
	}
	ret += "});\n"
	return ret
}
