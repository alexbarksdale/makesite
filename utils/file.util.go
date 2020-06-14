package utils

import (
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

func ReadFile(file string) string {
	fileContents, err := ioutil.ReadFile("corpus/" + file)
	if err != nil {
		log.Fatalf("Failed to read file! %v", err)
	}

	return string(fileContents)
}

func WriteFile(tmpl []byte, file string) {
	if err := ioutil.WriteFile("exports/"+file, tmpl, 0666); err != nil {
		log.Fatalf("Failed to write file %v", err)
	}
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)

	d := path.Dir(b)
	return filepath.Dir(d)

}
