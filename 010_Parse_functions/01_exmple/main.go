package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.New("Something").Parse("Here is the text in the template."))
	err := tpl.ExecuteTemplate(os.Stdout, "Something", nil)
	if err != nil {
		log.Fatal(err)
	}
}
