package utils

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var fileDesc = "Provide the path to the file to convert."

func Flags() string {
	var file string

	flag.StringVar(&file, "file", "", fileDesc)
	flag.Parse()

	// This probably isn't the best way to implement this. Redo this later after further dev.
	// If this gets messy after more commands, create a util to display this.
	if file == "" {
		fmt.Println("Please provide a valid flag. \n\nOptions:")
		fmt.Printf("--file     %v", fileDesc)
		os.Exit(1)
	}

	// FIX: Probably will cause a bug when we add more commands... maybe
	if !strings.Contains(file, ".txt") {
		return "You must provide a valid .txt file. Use the flag '--file' for help."
	}

	return file
}
