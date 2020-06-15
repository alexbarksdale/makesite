package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)

	d := path.Dir(b)
	return filepath.Dir(d)

}

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

func SearchTxtFiles(directory string) []string {
	var files []string

	dir := RootDir() + "/" + directory
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// Prevent scanning subdirectories
		if info.IsDir() {
			return nil
		}

		fileName := info.Name()
		if strings.Contains(fileName, ".txt") {
			files = append(files, fileName)
		}

		return nil
	}); err != nil {
		panic(err)
	}

	return files
}
