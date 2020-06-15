package generate

import (
	"bytes"
	"html/template"
	"log"
	"makesite/utils"
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
	tmpl := GenContent(file)
	utils.WriteFile(tmpl, file)
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
