package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type sage struct {
	Name  string
	Motto string
}

func main() {
	sages := []sage{
		sage{Name: "Buddha",
			Motto: "The Belief of no  beliefs"},
		sage{Name: "Gandhi",
			Motto: "Be the change"},
		sage{Name: "Muhammad",
			Motto: "To overcome evil with good is good, to resist evil by evil is evil"},
	}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatal(err)
	}

}
