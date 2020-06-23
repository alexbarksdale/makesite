package generate

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"makesite/utils"

	"gopkg.in/russross/blackfriday.v2"
)

type tmplData struct {
	Content string
}

func GenContent(filename string) []byte {
	var tmplPath string = "templates/template.tmpl"

	tmpl, err := template.New("template.tmpl").ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("Failed to create a new template! %v", err)
	}

	contentValue := tmplData{utils.ReadFile(filename)}
	var b bytes.Buffer
	if err := tmpl.Execute(&b, contentValue); err != nil {
		log.Fatalf("Failed to execute the template! %v", err)
	}
	return b.Bytes()
}

func GenSite(file string) {
	if utils.GetExtension(file) == ".md" {
		fmt.Printf("Converting %v file to html \n", file)
		tmpl := GenContent(file)
		output := blackfriday.Run(tmpl)
		utils.WriteFile(output, file)
	} else {
		fmt.Printf("Converting %v file to html \n", file)
		tmpl := GenContent(file)
		utils.WriteFile(tmpl, file)
	}
}

func GenDirContent(directory string) {
	validDirectory := utils.DoesDirectoryExist(directory)
	if !(validDirectory) {
		log.Fatalf("Directory '%v' does not exist!", directory)
	}

	files := utils.SearchTxtFiles(directory)
	for _, file := range files {
		GenSite(file)
	}
}
