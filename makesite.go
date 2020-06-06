package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type tmplData struct {
	Content string
}

func readFile() string {
	fileContents, err := ioutil.ReadFile("corpus/first-post.txt")
	if err != nil {
		log.Fatalf("Failed to read file! %v", err)
	}

	return string(fileContents)
}

func renderTemplate() {
	var tmplPath string = "/Users/alex/dev/backend/makesite/templates/template.tmpl"
	contentValue := tmplData{readFile()}

	tmpl, err := template.New("template.tmpl").ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("Failed to create a new template! %v", err)
	}
	if err := tmpl.Execute(os.Stdout, contentValue); err != nil {
		log.Fatalf("Failed to execute the template! %v", err)
	}
}

func main() {
	fmt.Printf("Make Site | A simple cli tool to generate a static site.")
	renderTemplate()
}
