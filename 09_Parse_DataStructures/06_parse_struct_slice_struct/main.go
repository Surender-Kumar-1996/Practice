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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func main() {
	sages := []sage{
		sage{
			Name:  "Buddha",
			Motto: "Healthy body leads healthy minds",
		},
		sage{
			Name:  "Rabindra nath Tagore",
			Motto: "You give me blood, i'll give you freedom",
		},
		sage{
			Name:  "Volkswagon",
			Motto: "Thus Auto",
		},
	}

	cars := []car{
		car{
			Manufacturer: "Suzuki",
			Model:        "Baleno",
			Doors:        4,
		},
		car{
			Manufacturer: "Mahindra",
			Model:        "Thar",
			Doors:        2,
		},
		car{
			Manufacturer: "Maruti",
			Model:        "Omini",
			Doors:        3,
		},
	}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatal(err)
	}
}
