package main

import (
	"fmt"
	"makesite/generate"
	"makesite/utils"
)

func main() {
	fmt.Printf("Make Site | A simple cli tool to generate a static site. \n")

	// temporary
	utils.WriteFile(generate.GenContent())
}
