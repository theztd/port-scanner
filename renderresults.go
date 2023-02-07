package main

import (
	"html/template"
	"os"
)

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
	t.ParseGlob("templates/*.html")

	t.ExecuteTemplate(os.Stdout, "statusPage.html", data)
}
