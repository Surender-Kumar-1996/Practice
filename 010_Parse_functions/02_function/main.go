package main

import (
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(str string) string {
	str = strings.TrimSpace(str)
	str = str[:3]
	return str
}

func main() {

}
