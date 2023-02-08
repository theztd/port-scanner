package main

import (
	"os"
	"text/template"
)

var templateFile string = "statusPage.htm"

func renderResults(data []hostStatus) {
	t := template.New("").Funcs(template.FuncMap{
		"isUp": func(str string) string {
			if str == "Open" {
				return "up"
			} else {
				return "down"
			}
		},
	})
	t.ParseGlob("templates/*")

	t.ExecuteTemplate(os.Stdout, templateFile, data)
}
