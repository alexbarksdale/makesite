package main

import (
	"fmt"
	"makesite/generate"
	"makesite/utils"
)

func main() {
	fmt.Printf("Make Site | A simple cli tool to generate a static site. \n\n")

	file, dir := utils.Flags()
	if dir != "" {
		generate.GenDirContent(dir)
	}
	if file != "" {
		generate.GenSite(file)
	}

}
