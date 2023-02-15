package main

import (
	"embed"
	"os"
	"text/template"
)

//go:embed templates
var templatesFS embed.FS

func renderResults(data []hostStatus, templateName, customTemplate string, out *os.File) {

	// Init demplate
	t := template.New("").Funcs(template.FuncMap{
		"isUp": func(str string) string {
			if str == "Open" {
				return "up"
			} else {
				return "down"
			}
		},
	})

	// t.ParseGlob("templates/*")
	t.ParseFS(templatesFS, "templates/*")

	// use custom template
	if customTemplate != "" {
		t.ParseFiles(customTemplate)
		t.ExecuteTemplate(out, templateName, data)
	} else {
		t.ExecuteTemplate(out, templateName, data)
	}
}
