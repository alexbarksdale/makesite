package utils

import (
	"io/ioutil"
	"log"
)

func ReadFile() string {
	fileContents, err := ioutil.ReadFile("corpus/first-post.txt")
	if err != nil {
		log.Fatalf("Failed to read file! %v", err)
	}

	return string(fileContents)
}

func WriteFile(tmpl []byte) {
	if err := ioutil.WriteFile("exports/first-post.html", tmpl, 0666); err != nil {
		log.Fatalf("Failed to write file %v", err)
	}
}
