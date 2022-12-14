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
	buddha := sage{
		Name:  "Buddha",
		Motto: "The Belief of no  beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)

	if err != nil {
		log.Fatal(err)
	}

}
