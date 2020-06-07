package main

import (
	"fmt"
	"makesite/generate"
	"makesite/utils"
)

func main() {
	fmt.Printf("Make Site | A simple cli tool to generate a static site. \n\n")

	file := utils.Flags()
	if file != "" {
		// TODO: Create helper to do this
		b := generate.GenContent(file)
		utils.WriteFile(b, file)
	}
}
