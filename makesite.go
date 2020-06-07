package main

import (
	"fmt"
	"makesite/utils"
)

func main() {
	fmt.Printf("Make Site | A simple cli tool to generate a static site. \n")

	// utils.WriteFile(generate.GenContent())
	file := utils.Flags()
	fmt.Println(file)
}
