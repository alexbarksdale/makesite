package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() string {
	fileContents, err := ioutil.ReadFile("corpus/first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate() {}

func main() {
	fmt.Println("Make Site | A simple cli tool to generate a static site.")
	output := readFile()
	fmt.Println(output)
}
